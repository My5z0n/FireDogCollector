package data

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
)

type TraceRepository struct {
	DB driver.Conn
}

func (m TraceRepository) GetTracesWithAnomalies(offset int, limit int) (result []dto.Trace) {

	query := `SELECT subq.*, predictions.anomaly_detected FROM (SELECT trace_id, start_time  FROM traces
		ORDER BY start_time DESC
		LIMIT ? OFFSET ?) as subq
    	LEFT JOIN predictions m on subq.trace_id = m.trace_id`
	result = []dto.Trace{}

	err := m.DB.Select(context.Background(), &result, query, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	println(len(result))
	return
}
