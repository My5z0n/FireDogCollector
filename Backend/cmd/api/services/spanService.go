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

	isParent := make(map[string]int)

	for _, span := range spans {
		if span.ParentSpanID != "" {
			isParent[span.ParentSpanID] = 1
		}
	}
	for _, span := range spans {
		if _, ok := isParent[span.SpanID]; !ok {
			spans = append(spans, dto.SpanListElementDTO{
				TraceID:      span.TraceID,
				SpanName:     "!END",
				ParentSpanID: span.SpanID,
				StartTime:    span.EndTime,
				EndTime:      span.EndTime,
			})
		}
	}

	prediction, err := s.Models.PredictionsRepository.GetAnomalyFromTraceID(id)
	if err != nil {
		//Not Found
		if err.Error() == "sql: no rows in result set" {
			return spans, nil
		} else {
			return nil, err
		}
		//

	}
	prediction.FitToSpans(spans)

	return spans, nil

}
