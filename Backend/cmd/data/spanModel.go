package data

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

type Span struct {
	TraceID      string    `json:"trace_id"`
	SpanID       string    `json:"span_id"`
	ParentSpanID string    `json:"parent_span_id"`
	SpanName     string    `json:"span_name"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Attributes   string    `json:"attributes"`
}

type SpanModel struct {
	DB driver.Conn
}

func (m SpanModel) Get(spanID string) {

}
