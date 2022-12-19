package data

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/models"
)

type PredictionsRepository struct {
	DB driver.Conn
}

func (r PredictionsRepository) GetAnomalyFromTraceID(trace_id string) (*models.Predictions, error) {

	query := `SELECT * from predictions WHERE trace_id = ?`

	row, err := r.DB.Query(context.Background(), query, trace_id)
	if err != nil {

		return nil, err
	}
	anomaly := models.Predictions{}
	err = row.ScanStruct(&anomaly)
	if err != nil {
		return nil, err
	}

	return &anomaly, nil

}
