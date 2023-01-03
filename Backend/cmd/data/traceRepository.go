package data

import (
	"context"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/models"
)

type TraceRepository struct {
	DB driver.Conn
}

func (m TraceRepository) GetTracesWithAnomalies(offset int, limit int) (result []dto.TracesListElement) {

	query := `SELECT subq.*, predictions.anomaly_detected FROM (SELECT trace_id, start_time  FROM traces
		ORDER BY start_time DESC
		LIMIT ? OFFSET ?) as subq
    	LEFT JOIN predictions m on subq.trace_id = m.trace_id`
	result = []dto.TracesListElement{}

	err := m.DB.Select(context.Background(), &result, query, limit, offset)
	if err != nil {
		fmt.Println(err)
	}
	println(len(result))
	return
}

func (m TraceRepository) GetSingleTrace(traceID string) (*models.Trace, error) {

	query := `select * from traces WHERE trace_id = ?;`

	var result models.Trace

	row := m.DB.QueryRow(context.Background(), query, traceID)
	if err := row.ScanStruct(&result); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil

}
