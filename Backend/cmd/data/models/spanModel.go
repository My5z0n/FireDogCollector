package models

import "time"

type Span struct {
	TraceID      string         `json:"trace_id"`
	SpanID       string         `json:"span_id"`
	ParentSpanID string         `json:"parent_span_id"`
	SpanName     string         `json:"span_name"`
	StartTime    time.Time      `json:"start_time"`
	EndTime      time.Time      `json:"end_time"`
	Attributes   map[string]any `json:"attributes"`
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
