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
	tracesRoute.GET("/:traceid", s.Controllers.TraceControllers.GetOne)

	spanRoute := s.RouterGroup.Group("/spans")
	spanRoute.GET("/:spanid", s.Controllers.SpanControllers.GetOne)

	anomalyRoute := s.RouterGroup.Group("/anomalydetetor")
	anomalyRoute.GET("/starttrain", s.Controllers.AnomalyDetectorControllers.InitLearningModel)
	anomalyRoute.POST("/find-outlines", s.Controllers.AnomalyDetectorControllers.HuntOutlines)

}
