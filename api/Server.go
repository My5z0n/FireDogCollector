package api

import (
	"context"
	"fmt"
	"github.com/My5z0n/FireDogCollector/models"
	"github.com/My5z0n/FireDogCollector/repository"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"go.opentelemetry.io/proto/otlp/trace/v1"
	"time"
)

type Server struct {
	coltracepb.UnimplementedTraceServiceServer
	Abba            string
	TraceRepository repository.TraceRepository
}

// SayHello implements helloworld.GreeterServer
func (s *Server) Export(ctx context.Context, request *coltracepb.ExportTraceServiceRequest) (*coltracepb.ExportTraceServiceResponse, error) {

	sp := request.ResourceSpans

	for _, resSpan := range sp {
		for _, scopeSpan := range resSpan.GetScopeSpans() {
			for _, span := range scopeSpan.GetSpans() {
				name := span.Name
				fmt.Printf("Name is: %s \n", name)
				go s.processSpan(resSpan, scopeSpan, span)
			}
		}
	}
	//fmt.Printf("Got:", sp[0].String())
	return &coltracepb.ExportTraceServiceResponse{}, nil
}

func (s *Server) processSpan(resource *v1.ResourceSpans, scope *v1.ScopeSpans, span *v1.Span) {
	err := s.TraceRepository.SaveSpan(models.SaveSpan{
		Trace_id:       string(span.TraceId[:]),
		Span_id:        string(span.SpanId[:]),
		Parent_span_id: string(span.ParentSpanId[:]),
		Span_name:      span.Name,
		Start_time:     time.Unix(0, int64(span.StartTimeUnixNano)),
	},
	)
}
