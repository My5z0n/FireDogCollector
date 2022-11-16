package dto

import "time"

type Trace struct {
	TraceID   string    `ch:"trace_id"`
	StartTime time.Time `ch:"start_time"`
}
