package data

import "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

type Repositories struct {
	SpanRepository  SpanRepository
	TraceRepository TraceRepository
	//Predictions PredictModel
}

func NewModels(db driver.Conn) Repositories {
	return Repositories{SpanRepository: SpanRepository{DB: db},
		TraceRepository: TraceRepository{DB: db}}
}
