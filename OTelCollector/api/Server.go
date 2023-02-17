package api

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/My5z0n/FireDogCollector/OtelCollector/models"
	"github.com/My5z0n/FireDogCollector/OtelCollector/repository"
	"github.com/My5z0n/FireDogCollector/OtelCollector/utils"
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
	TraceRepository repository.TraceRepository
}

func (s *Server) processSpan(resource *v1.ResourceSpans, scope *v1.ScopeSpans, span *v1.Span, processedSpanResultChan chan<- models.ClickHouseSpan) error {

	attributesMap := make(map[string]any)
	externalParamNames := make([]string, 0, 10)
	for i, v := range span.Attributes {
		if utils.CheckColNames(v.Key) {
			switch k := span.Attributes[i].GetValue().GetValue().(type) {
			case *commonpb.AnyValue_StringValue:
				attributesMap[v.Key] = k.StringValue
				externalParamNames = append(externalParamNames, v.Key)
			case *commonpb.AnyValue_BoolValue:
				attributesMap[v.Key] = k.BoolValue
				externalParamNames = append(externalParamNames, v.Key)
			case *commonpb.AnyValue_IntValue:
				attributesMap[v.Key] = k.IntValue
				externalParamNames = append(externalParamNames, v.Key)
			case *commonpb.AnyValue_DoubleValue:
				attributesMap[v.Key] = k.DoubleValue
				externalParamNames = append(externalParamNames, v.Key)
			case *commonpb.AnyValue_BytesValue:
				attributesMap[v.Key] = hex.EncodeToString(k.BytesValue)
				externalParamNames = append(externalParamNames, v.Key)
			case nil:
			default:
			}
		}

	}

	model := models.ClickHouseSpan{
		Trace_id:         hex.EncodeToString(span.TraceId),
		Span_id:          hex.EncodeToString(span.SpanId),
		Parent_span_id:   hex.EncodeToString(span.ParentSpanId),
		Span_name:        span.Name,
		Start_time:       time.Unix(0, int64(span.StartTimeUnixNano)),
		End_time:         time.Unix(0, int64(span.EndTimeUnixNano)),
		ExternalColNames: externalParamNames,
		Attributes:       attributesMap,
	}
	processedSpanResultChan <- model
	err := s.TraceRepository.SaveSpan(model)

	return err
}

func (s *Server) saveTrace(inputProcessedSpansChan <-chan models.ClickHouseSpan, counter int32) error {

	var startTime time.Time
	var currentTraceID string
	var spanRoot string
	spanLeafList := []string{}

	spansMap := make(map[string]*models.Span)

	for counter > 0 {
		//Load next Span
		v := <-inputProcessedSpansChan
		counter -= 1

		vSpan := models.SpanAttributes{
			Span_ID:        v.Span_id,
			Span_Name:      v.Span_name,
			Parent_Span_id: v.Parent_span_id,
			Start_time:     v.Start_time,
			End_time:       v.End_time,
		}
		if vSpan.Parent_Span_id == "" {
			spanRoot = v.Span_id
		}

		//Check traceID
		if currentTraceID == "" {
			currentTraceID = v.Trace_id
			startTime = v.Start_time
		} else if currentTraceID != v.Trace_id {
			log.Printf("Foreign id detected: Expected %s got %s skiping...\n", currentTraceID, v.Trace_id)
			continue
		}

		//Register Span
		if _, ok := spansMap[v.Span_id]; !ok {
			spansMap[v.Span_id] = &models.Span{
				SpanProperties: &vSpan,
				SpanChildren:   make([]*models.Span, 0),
			}
		} else {
			spansMap[v.Span_id].SpanProperties = &vSpan
		}

		//AddRelation
		if _, ok := spansMap[v.Parent_span_id]; ok {
			spansMap[v.Parent_span_id].SpanChildren = append(spansMap[v.Parent_span_id].SpanChildren, spansMap[v.Span_id])
		} else if v.Parent_span_id != "" {
			spansMap[v.Parent_span_id] = &models.Span{SpanChildren: []*models.Span{spansMap[v.Span_id]}}
		}
	}

	//Get span leafs
	for k, v := range spansMap {
		if len(v.SpanChildren) == 0 {
			spanLeafList = append(spanLeafList, k)
		}
	}

	paths := spanPathsProcessing.GeneratePathsFromSpans(spansMap, spanLeafList)
	str, _ := json.Marshal(spansMap[spanRoot])
	err := s.TraceRepository.SaveTrace(paths, currentTraceID, startTime, string(str))
	return err
}

func (s *Server) Export(ctx context.Context, request *coltracepb.ExportTraceServiceRequest) (*coltracepb.ExportTraceServiceResponse, error) {

	g := new(errgroup.Group)

	//Maximum 1000 of elements
	processSpanChan := make(chan models.ClickHouseSpan, 1000)

	var waitCount int32

	for _, resSpan := range request.ResourceSpans {
		for _, scopeSpan := range resSpan.GetScopeSpans() {
			for _, span := range scopeSpan.GetSpans() {
				name := span.Name
				fmt.Printf("Start processing [Name]: %s \n", name)

				atomic.AddInt32(&waitCount, 1)
				g.Go(func() error {
					err := s.processSpan(resSpan, scopeSpan, span, processSpanChan)
					return err

				})

			}
		}
	}

	if err := g.Wait(); err != nil {
		log.Printf("Error occured during SpanExport %v \n", err)
		return &coltracepb.ExportTraceServiceResponse{}, err
	}

	err := s.saveTrace(processSpanChan, waitCount)
	if err != nil {
		log.Printf("Error occured during saveTrace %v \n", err)
		return &coltracepb.ExportTraceServiceResponse{}, err
	}

	return &coltracepb.ExportTraceServiceResponse{}, nil
}
