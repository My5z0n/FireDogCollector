package controllers

import "github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"

type Controllers struct {
	SpanControllers  SpanController
	TraceControllers TraceController
}

func NewControllers(s services.Services) Controllers {
	return Controllers{SpanControllers: SpanController{Services: s},
		TraceControllers: TraceController{Services: s},
	}
}
