package data

import "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

type Models struct {
	SpanModel  SpanModel
	TraceModel TraceModel
	//Predictions PredictModel
}

func NewModels(db driver.Conn) Models {
	return Models{SpanModel: SpanModel{DB: db},
		TraceModel: TraceModel{DB: db}}
}
