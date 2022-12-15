package data

import (
	"context"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/models"
	"reflect"
)

type SpanRepository struct {
	DB driver.Conn
}

func (m SpanRepository) GetSpan(spanID string) (*models.Span, error) {

	query := `SELECT * from spans WHERE span_id = ?`

	rows, err := m.DB.Query(context.Background(), query, spanID)
	if err != nil {
		return nil, err
	}
	var columnTypes = rows.ColumnTypes()
	var columnNames = rows.Columns()
	var additionalAttributes = make([]interface{}, len(columnTypes))

	for i := range columnTypes {
		additionalAttributes[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}

	rows.Next()
	if err := rows.Scan(additionalAttributes...); err != nil {
		return nil, err
	}

	var ResultSpan = models.Span{Attributes: make(map[string]any)}

	for i, colName := range columnNames {
		ResultSpan.SetAttribute(colName, additionalAttributes[i])
	}

	return &ResultSpan, nil
}
