package services

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data/dto"
)

type SpanService struct {
	Models data.Repositories
}

func (s SpanService) GetSpan(id string) interface{} {

	result, err := s.Models.SpanRepository.GetSpan(id)
	if err != nil {
		return err
	}
	return result

}

func (s SpanService) GetSpansListFromTraceID(id string) ([]dto.SpanListElementDTO, error) {

	spans, err := s.Models.SpanRepository.GetSpansFromTraceID(id)
	if err != nil {
		return nil, err
	}
	prediction, err := s.Models.PredictionsRepository.GetAnomalyFromTraceID(id)
	if err != nil {
		//Not Found
		if err.Error() == "EOF" {
			return spans, nil
		} else {
			return nil, err
		}
		//

	}
	prediction.FitToSpans(spans)

	return spans, nil

}
