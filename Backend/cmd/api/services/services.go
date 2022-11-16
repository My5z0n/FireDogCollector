package services

import "github.com/My5z0n/FireDogCollector/Backend/cmd/data"

type Services struct {
	SpanService  SpanService
	TraceService TraceService
}

func NewServices(models data.Models) Services {
	return Services{SpanService: SpanService{
		Models: models},
		TraceService: TraceService{
			Models: models},
	}
}
