package server

import "github.com/gin-gonic/gin"

type BackendRouter struct {
	router *gin.RouterGroup
}

func (r *BackendRouter) Attach() {

	tracesRoute := r.router.Group("/traces")
	tracesRoute.GET("/")

}

func NewBackendRouter(e *gin.Engine) *BackendRouter {
	mainRouter := BackendRouter{
		router: e.Group("/api/v1"),
	}
	mainRouter.Attach()
	return &BackendRouter{}
}
