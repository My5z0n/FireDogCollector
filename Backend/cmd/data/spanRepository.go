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

func (m SpanRepository) GetSpan(spanID string) interface{} {

	query := `SELECT * from spans WHERE span_id = ?`
	//result = []dto.TracesListElement{}

	rows, err := m.DB.Query(context.Background(), query, spanID)
	if err != nil {
		return err
	}
	var columnTypes = rows.ColumnTypes()
	var columnNames = rows.Columns()
	var vars = make([]interface{}, len(columnTypes))

	var ResultSpan = models.Span{Attributes: make(map[string]any)}

	for i := range columnTypes {
		vars[i] = reflect.New(columnTypes[i].ScanType()).Interface()
	}

	rows.Next()
	if err := rows.Scan(vars...); err != nil {
		return err
	}
	for i, v := range columnNames {
		ResultSpan.SetAttribute(v, vars[i])
	}

	return ResultSpan.Attributes
}
