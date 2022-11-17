package data

import (
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type SpanRepository struct {
	DB driver.Conn
}

func (m SpanRepository) Get(spanID string) {

}
