package services

import "github.com/My5z0n/FireDogCollector/Backend/cmd/data"

type Services struct {
	SpanService    SpanService
	TraceService   TraceService
	AnomalyService AnomalyService
}

func NewServices(models data.Repositories) Services {
	return Services{SpanService: SpanService{
		Models: models},
		TraceService: TraceService{
			Models: models},
		AnomalyService: AnomalyService{
			Models: models,
		},
	}
}
