package services

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
)

type TraceService struct {
	Models data.Models
}

func (s TraceService) GetTraces(page int) []dto.Trace {
	PageSize := 2

	result := s.Models.TraceModel.GetTracesWithAnomalies(page*PageSize, PageSize)
	return result

}
