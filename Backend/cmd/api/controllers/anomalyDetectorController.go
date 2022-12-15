package controllers

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnomalyDetectorController struct {
	Services services.Services
}

// GetOne godoc
// @Summary Get specific Span
// @Schemes
// @Description Return specific span
// @Tags Span
// @Produce json
// @Param spanid path string true "Span ID"
// @Router /spans/{spanid} [get]
// @host localhost:9900
func (c AnomalyDetectorController) InitLearningModel(ctx *gin.Context) {

	resp, err := http.Get("http://localhost/9181/START_TRAIN")
	if err != nil {
		//TODO: HANDLE ERROR
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c AnomalyDetectorController) StartRestartModelLearning(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "OK")
}
