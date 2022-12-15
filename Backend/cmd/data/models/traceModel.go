package models

import "time"

type PathArrayElement struct {
	SpanName string `ch:"span_name"`
	SpanID   string `ch:"span_id"`
}

type Trace struct {
	TraceID    string               `ch:"trace_id"`
	Paths      string               `ch:"paths"`
	PathsArray [][]PathArrayElement `ch:"paths_array"`
	StartTime  time.Time            `ch:"start_time"`
	JsonSpans  string               `ch:"json_spans"`
}
