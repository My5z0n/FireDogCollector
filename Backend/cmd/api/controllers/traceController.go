package controllers

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TraceController struct {
	Services services.Services
}

type TraceControllerBinding struct {
	Page int `form:"page"`
}

func (c TraceController) Get(ctx *gin.Context) {
	binding := TraceControllerBinding{}

	ctx.ShouldBindQuery(&binding)

	res := c.Services.TraceService.GetTraces(binding.Page)

	ctx.JSON(http.StatusOK, res)
}
