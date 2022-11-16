package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Sever struct {
	Engine        *gin.Engine
	BackendRouter *BackendRouter
}

func (s *Sever) Serve() {
	err := s.Engine.Run("localhost:9900")
	if err != nil {
		log.Panicf("Eror During Start %v", err)
	}
}

func CreateNew(engine *gin.Engine, backendRouter *BackendRouter) *Sever {
	return &Sever{Engine: engine, BackendRouter: backendRouter}
}
