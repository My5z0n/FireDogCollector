package services

import (
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

	result := dto.TraceModelDTO{
		TraceID:   trace.TraceID,
		StartTime: trace.StartTime,
	}

	return result
}
