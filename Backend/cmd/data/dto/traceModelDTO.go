package dto

import (
	"time"
)

type TraceModelDTO struct {
	TraceID   string
	StartTime time.Time
	Anomaly   bool
	SpansList []SpanListElementDTO
	//SpansMap  map[string]any
}
