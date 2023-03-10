package dto

import "time"

type SpanListElementDTO struct {
	TraceID                 string    `json:"trace_id"`
	SpanID                  string    `json:"span_id"`
	ParentSpanID            string    `json:"parent_span_id"`
	SpanName                string    `json:"span_name"`
	StartTime               time.Time `json:"start_time"`
	EndTime                 time.Time `json:"end_time"`
	AnomalyDetected         bool      `json:"AnomalyDetected"`
	ExpectedAnomalySpanName string    `json:"ExpectedAnomalySpanName"`
	AnomalyPositionInTrace  int       `json:"AnomalyPositionInTrace"`
}
