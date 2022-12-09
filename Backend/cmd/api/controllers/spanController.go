package controllers

import (
	"fmt"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SpanController struct {
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
func (c SpanController) GetOne(ctx *gin.Context) {
	binding := struct {
		SpanID string `uri:"spanid" binding:"required"`
	}{}

	err := ctx.ShouldBindUri(&binding)
	if err != nil {
		//TODO: PROPERLY HANDLE ERRORS
		fmt.Printf("ERR: %v", err)
		return
	}

	res := c.Services.SpanService.GetSpan(binding.SpanID)

	ctx.JSON(http.StatusOK, res)

}
