package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/My5z0n/FireDogCollector/OtelCollector/models"
	"github.com/My5z0n/FireDogCollector/OtelCollector/repository"
	"github.com/My5z0n/FireDogCollector/OtelCollector/utils/spanPathsProcessing"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	"go.opentelemetry.io/proto/otlp/trace/v1"
	"golang.org/x/sync/errgroup"
	"log"
	"sync/atomic"
	"time"
)

type Server struct {
	coltracepb.UnimplementedTraceServiceServer
	Abba            string
	TraceRepository repository.TraceRepository
}

func (s *Server) processSpan(resource *v1.ResourceSpans, scope *v1.ScopeSpans, span *v1.Span, processedSpanResultChan chan<- models.ClickHouseSpan) error {

	attributesMap := make(map[string]any)
	for i, v := range span.Attributes {
		switch k := span.Attributes[i].GetValue().GetValue().(type) {
		case *commonpb.AnyValue_StringValue:
			attributesMap[v.Key] = k.StringValue
		case *commonpb.AnyValue_BoolValue:
			attributesMap[v.Key] = k.BoolValue
		case *commonpb.AnyValue_IntValue:
			attributesMap[v.Key] = k.IntValue
		case *commonpb.AnyValue_DoubleValue:
			attributesMap[v.Key] = k.DoubleValue
		case *commonpb.AnyValue_BytesValue:
			attributesMap[v.Key] = hex.EncodeToString(k.BytesValue)
		case nil:
		default:
		}
	}

	//obj, errr := json.Marshal(attributesMap)
	//fmt.Println("Marshal Datasets Result : ", string(obj), errr)

	model := models.ClickHouseSpan{
		Trace_id:       hex.EncodeToString(span.TraceId),
		Span_id:        hex.EncodeToString(span.SpanId),
		Parent_span_id: hex.EncodeToString(span.ParentSpanId),
		Span_name:      span.Name,
		Start_time:     time.Unix(0, int64(span.StartTimeUnixNano)),
		End_time:       time.Unix(0, int64(span.EndTimeUnixNano)),
		Attributes:     attributesMap,
	}
	processedSpanResultChan <- model
	err := s.TraceRepository.SaveSpan(model)

	return err
}

func (s *Server) saveTrace(dataDigChan <-chan models.ClickHouseSpan, counter int32) error {

	spanDependencyMap := make(map[string]models.SpanTag)
	spanChildCounter := make(map[string]int)
	spanLeafList := []string{}
	currentTraceID := ""
	var startTime time.Time

	for counter > 0 {
		//Load next Span
		v := <-dataDigChan
		counter -= 1

		//Check traceID
		if currentTraceID == "" {
			currentTraceID = v.Trace_id
			startTime = v.Start_time
		} else if currentTraceID != v.Trace_id {
			log.Printf("Foreign id detected: Expected %s got %s skiping...\n", currentTraceID, v.Trace_id)
			continue
		}

		//Count child spans
		spanChildCounter[v.Parent_span_id] = spanChildCounter[v.Parent_span_id] + 1
		if _, ok := spanChildCounter[v.Span_id]; !ok {
			spanChildCounter[v.Span_id] = 0
		}

		//Save span node
		spanDependencyMap[v.Span_id] = models.SpanTag{
			Start_time:     v.Start_time,
			Span_Name:      v.Span_name,
			Parent_Span_id: v.Parent_span_id,
		}

	}

	//Get span leafs
	for k, v := range spanChildCounter {
		if v == 0 {
			spanLeafList = append(spanLeafList, k)
		}
	}

	paths := spanPathsProcessing.GeneratePathsFromSpans(spanDependencyMap, spanLeafList)
	err := s.TraceRepository.SaveTrace(paths, currentTraceID, startTime)
	return err
}

func (s *Server) Export(ctx context.Context, request *coltracepb.ExportTraceServiceRequest) (*coltracepb.ExportTraceServiceResponse, error) {

	sp := request.ResourceSpans

	g := new(errgroup.Group)

	//Maximum 1000 of elements
	processSpanChan := make(chan models.ClickHouseSpan, 1000)

	var waitCount int32

	for _, resSpan := range sp {
		for _, scopeSpan := range resSpan.GetScopeSpans() {
			for _, span := range scopeSpan.GetSpans() {
				name := span.Name
				fmt.Printf("Start processing [Name]: %s \n", name)

				atomic.AddInt32(&waitCount, 1)
				g.Go(func() error {
					err := s.processSpan(resSpan, scopeSpan, span, processSpanChan)
					//if err != nil {
					//	atomic.AddInt32(&waitCount, -1)
					//}
					return err

				})

			}
		}
	}
	/*
		if err := g.Wait(); err != nil {
			log.Printf("Error occured during SpanExport %v \n", err)
			return &coltracepb.ExportTraceServiceResponse{}, err
		}*/

	err := s.saveTrace(processSpanChan, waitCount)
	if err != nil {
		log.Printf("Error occured during saveTrace %v \n", err)
		return &coltracepb.ExportTraceServiceResponse{}, err
	}

	return &coltracepb.ExportTraceServiceResponse{}, nil
}
