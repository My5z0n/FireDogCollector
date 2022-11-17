package server

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/controllers"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	"github.com/gin-gonic/gin"
)

type Sever struct {
	Engine      *gin.Engine
	RouterGroup *gin.RouterGroup
	Models      data.Repositories
	Services    services.Services
	Controllers controllers.Controllers
}

func (s Sever) AttachRoutes() {
	s.AttachControllers()
	s.AttachHealthCheck()
}
func CreateNew(engine *gin.Engine, m data.Repositories, srcs services.Services, c controllers.Controllers) (server *Sever) {
	server = &Sever{
		Engine:      engine,
		RouterGroup: engine.Group("/api/v1"),
		Models:      m,
		Services:    srcs,
		Controllers: c,
	}
	server.AttachRoutes()

	return
}
func (s Sever) Serve() error {
	err := s.Engine.Run("localhost:9900")
	return err
}
