package models

import "time"

type ClickHouseSpan struct {
	Trace_id       string         `ch:"trace_id"`
	Span_id        string         `ch:"span_id"`
	Parent_span_id string         `ch:"parent_span_id"`
	Span_name      string         `ch:"span_name"`
	Start_time     time.Time      `ch:"start_time"`
	End_time       time.Time      `ch:"end_time"`
	Attributes     map[string]any `ch:"attributes"`
}

type SpanTag struct {
	Start_time     time.Time
	Span_Name      string
	Parent_Span_id string
}
