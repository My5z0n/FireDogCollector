package services

import "github.com/My5z0n/FireDogCollector/Backend/cmd/data"

type TraceService struct {
	Models data.Models
}

func (s TraceService) GetTraces(page int) []data.Trace {
	PageSize := 2

	result := s.Models.TraceModel.GetTraces(page*PageSize, PageSize)
	return result

}
