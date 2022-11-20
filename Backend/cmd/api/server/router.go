package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s Sever) AttachHealthCheck() {

	s.RouterGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}
func (s Sever) AttachControllers() {

	tracesRoute := s.RouterGroup.Group("/traces")
	tracesRoute.GET("/", s.Controllers.TraceControllers.GetAll)
	tracesRoute.GET("/:traceid/", s.Controllers.TraceControllers.GetOne)

}
