package controllers

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type TraceController struct {
	Services services.Services
}

type TraceControllerBinding struct {
	Page int `form:"page" binding:"numeric,gte=0"`
}

// Get godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Param page query int false "page"
// @Router /traces [get]
// @host localhost:9900
func (c TraceController) Get(ctx *gin.Context) {
	binding := TraceControllerBinding{}

	err := ctx.ShouldBindQuery(&binding)
	if err != nil {
		//TODO: PROPERLY HANDLE ERRORS
		log.Error().Msg("BAD INPUT")
		return
	}

	res := c.Services.TraceService.GetTraces(binding.Page)

	ctx.JSON(http.StatusOK, res)
}
