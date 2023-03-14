package models

import (
	"database/sql"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
)

type Predictions struct {
	TraceID                string       `ch:"trace_id"`
	AnomalyDetected        sql.NullBool `ch:"anomaly_detected"`
	SpanName               string       `ch:"span_name"`
	SpanID                 string       `ch:"span_id"`
	ExpectedSpanName       string       `ch:"expected_span_name"`
	AnomalyPositionInTrace *int32       `ch:"span_position"`
}

func (p Predictions) FitToSpans(spanList []dto.SpanListElementDTO) {

	if p.AnomalyDetected.Valid != true || p.AnomalyDetected.Bool != true {
		return
	}

	for i := range spanList {
		if spanList[i].SpanID == p.SpanID {
			spanList[i].AnomalyDetected = p.AnomalyDetected.Bool
			spanList[i].ExpectedAnomalySpanName = p.ExpectedSpanName
			spanList[i].AnomalyPositionInTrace = int(*p.AnomalyPositionInTrace)
			break
		}
	}
}
