package services

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
	"strings"
)

type AnomalyService struct {
	Models data.Repositories
}

func (s AnomalyService) MakeOutlinesRequest(binding dto.SpanListElementDTO) (interface{}, error) {

	val, err := s.Models.TraceRepository.GetSingleTrace(binding.TraceID)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(val.Paths, "#")

	goodPart := parts[:binding.AnomalyPositionInTrace]
	goodPart = append(goodPart, binding.ExpectedAnomalySpanName)
	_ = parts[:binding.AnomalyPositionInTrace+1]

	return nil, nil
}
