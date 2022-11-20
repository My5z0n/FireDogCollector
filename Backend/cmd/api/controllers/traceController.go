package controllers

import (
	"fmt"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type TraceController struct {
	Services services.Services
}

// GetAll godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Router /traces [get]
// @host localhost:9900
func (c TraceController) GetAll(ctx *gin.Context) {
	type TraceControllerBinding struct {
		Page int `form:"page" binding:"numeric,gte=0"`
	}

	binding := TraceControllerBinding{}

	err := ctx.ShouldBindQuery(&binding)
	if err != nil {
		//TODO: PROPERLY HANDLE ERRORS
		log.Error().Msg("BAD INPUT")
		return
	}

	res := c.Services.TraceService.GetTracesWithAnomalies(binding.Page)

	ctx.JSON(http.StatusOK, res)
}

func (c TraceController) GetOne(ctx *gin.Context) {
	paramBinding := struct {
		TraceID string `uri:"traceid" binding:"required"`
	}{}

	err := ctx.ShouldBindUri(&paramBinding)
	if err != nil {
		//TODO: PROPERLY HANDLE ERRORS
		fmt.Printf("ERR: %v", err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"ret": paramBinding.TraceID})

}
