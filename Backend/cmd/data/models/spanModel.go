package models

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
	"time"
)

type Span struct {
	TraceID      string    `ch:"trace_id"`
	SpanID       string    `ch:"span_id"`
	ParentSpanID string    `ch:"parent_span_id"`
	SpanName     string    `ch:"span_name"`
	StartTime    time.Time `ch:"start_time"`
	EndTime      time.Time `ch:"end_time"`
	Attributes   map[string]any
}

func (s *Span) MakeDTO() dto.SpanListElementDTO {
	return dto.SpanListElementDTO{
		TraceID:                 s.TraceID,
		SpanID:                  s.SpanID,
		ParentSpanID:            s.ParentSpanID,
		SpanName:                s.SpanName,
		StartTime:               s.StartTime,
		EndTime:                 s.EndTime,
		AnomalyDetected:         false,
		ExpectedAnomalySpanName: "",
	}
}

func (s *Span) SetAttribute(name string, value any) {

	switch name {
	case "trace_id":
		s.TraceID = *(value).(*string)
	case "span_id":
		s.SpanID = *(value).(*string)
	case "parent_span_id":
		s.ParentSpanID = *(value).(*string)
	case "span_name":
		s.SpanName = *(value).(*string)
	case "start_time":
		s.StartTime = **(value).(**time.Time)
	case "end_time":
		s.EndTime = **(value).(**time.Time)
	default:
		switch v := value.(type) {
		case **string:
			if *v != nil {
				s.Attributes[name] = *v
			}
		case **int:
			if *v != nil {
				s.Attributes[name] = *v
			}
		case **float64:
			if *v != nil {
				s.Attributes[name] = *v
			}
		case **time.Time:
			if *v != nil {
				s.Attributes[name] = *v
			}

		}
	}
}
