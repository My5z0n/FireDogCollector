package models

import "database/sql"

type Predictions struct {
	TraceID          string       `ch:"trace_id"`
	AnomalyDetected  sql.NullBool `ch:"anomaly_detected"`
	SpanName         string       `ch:"span_name"`
	SpanID           string       `ch:"span_id"`
	ExpectedSpanName string       `ch:"expected_span_name"`
}
