package models

import "time"

type SaveSpan struct {
	Trace_id       string
	Span_id        string
	Parent_span_id string
	Span_name      string
	Start_time     time.Time
	End_time       time.Time
}
