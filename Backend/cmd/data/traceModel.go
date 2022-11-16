package data

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"time"
)

type PathArrayElement struct {
	SpanName string `ch:"span_name"`
	SpanID   string `ch:"span_id"`
}

type Trace struct {
	TraceID    string               `ch:"trace_id"`
	Paths      string               `ch:"paths"`
	PathsArray [][]PathArrayElement `ch:"paths_array"`
	StartTime  time.Time            `ch:"start_time"`
}

type TraceModel struct {
	DB driver.Conn
}

func (m TraceModel) GetTraces(offset int, limit int) (result []Trace) {
	query := `
		SELECT * FROM traces
		ORDER BY start_time DESC
		LIMIT ? OFFSET ?;
 		`
	result = []Trace{}

	err := m.DB.Select(context.Background(), &result, query, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	println(len(result))
	return
}
