package services

import "github.com/My5z0n/FireDogCollector/Backend/cmd/data"

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
