package services

import (
	"encoding/json"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
)

type TraceService struct {
	Models data.Repositories
}

// TODO: Handle Errors if err!= nil
func (s TraceService) GetTracesWithAnomalies(page int) []dto.TracesListElement {
	PageSize := 5

	result := s.Models.TraceRepository.GetTracesWithAnomalies(page*PageSize, PageSize)
	return result

}

func (s TraceService) GetSingleTraceWithAnomalyPrediction(traceID string) dto.TraceModelDTO {

	trace, err := s.Models.TraceRepository.GetSingleTrace(traceID)
	if trace == nil {
		print("Empty")
		return dto.TraceModelDTO{}
	}
	if err != nil {
		print(err)
		return dto.TraceModelDTO{}
	}

	anomaly, err := s.Models.TraceRepository.GetAnomaly(traceID)
	if err != nil {
		print(err)
		return dto.TraceModelDTO{}
	}
	jsonMap := make(map[string]interface{})

	json.Unmarshal([]byte(trace.JsonSpans), &jsonMap)

	result := dto.TraceModelDTO{
		TraceID:   trace.TraceID,
		StartTime: trace.StartTime,
		SpansMap:  jsonMap,
	}
	if anomaly != nil && anomaly.AnomalyDetected.Bool == true {
		result.ModifyAnomalySpans(anomaly.SpanID, anomaly.ExpectedSpanName)
	}

	return result
}
