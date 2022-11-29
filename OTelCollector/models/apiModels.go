package models

import "time"

type ClickHouseSpan struct {
	Trace_id         string         `ch:"trace_id"`
	Span_id          string         `ch:"span_id"`
	Parent_span_id   string         `ch:"parent_span_id"`
	Span_name        string         `ch:"span_name"`
	Start_time       time.Time      `ch:"start_time"`
	End_time         time.Time      `ch:"end_time"`
	Attributes       map[string]any `ch:"attributes"`
	ExternalColNames []string
}
type Span struct {
	SpanProperties *SpanAttributes
	SpanChildren   []*Span
}

type SpanAttributes struct {
	Span_ID            string
	Span_Name          string
	Parent_Span_id     string
	Start_time         time.Time
	End_time           time.Time
	Anomaly            bool
	Expected_Span_Name string
}
