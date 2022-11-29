package dto

import "time"

type TraceModelDTO struct {
	TraceID   string
	StartTime time.Time
	Anomaly   bool
	Spans     map[string]any
}

func (r *TraceModelDTO) ModifyAnomalySpans(spanID string, expectedSpanName string) {

	if dfsScan(r.Spans, spanID, expectedSpanName) {
		r.Anomaly = true
	} else {
		print("notFound")
	}

}

func dfsScan(spanMap map[string]any, searchedSpanID string, expectedSpanName string) bool {

	spanID := spanMap["SpanProperties"].(map[string]any)["Span_ID"].(string)
	childList := spanMap["SpanChildren"].([]interface{})

	if spanID == searchedSpanID {
		spanMap["SpanProperties"].(map[string]any)["Anomaly"] = true
		spanMap["SpanProperties"].(map[string]any)["Expected_Span_Name"] = expectedSpanName
		return true
	}

	if len(childList) != 0 {
		for _, v := range childList {
			if dfsScan(v.(map[string]any), searchedSpanID, expectedSpanName) {
				return true
			}
		}
		return false
	}
	return false

}
