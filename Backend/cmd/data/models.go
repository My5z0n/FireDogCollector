package data

import "github.com/ClickHouse/clickhouse-go/v2/lib/driver"

type Models struct {
	Spans       SpanModel
	Traces      TraceModel
	Predictions PredictModel
}

func getModels(db *driver.Conn) Models {
	return Models{}
}
