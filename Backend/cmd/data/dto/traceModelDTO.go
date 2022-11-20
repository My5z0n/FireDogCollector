package dto

import (
	"database/sql"
	"time"
)

type Trace struct {
	TraceID   string       `ch:"trace_id"`
	StartTime time.Time    `ch:"start_time"`
	Anomaly   sql.NullBool `ch:"anomaly_detected"`
}
