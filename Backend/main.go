package main

import (
	"fmt"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/controllers"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/server"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/services"
	"github.com/My5z0n/FireDogCollector/Backend/cmd/data"
	docs "github.com/My5z0n/FireDogCollector/Backend/docs"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"time"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.UnixDate})

	log.Info().Msg("Backend - FiredogTraces (c) 2022 - Szymon Nagel \n")
	ginEngine := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//gin.SetMode(gin.ReleaseMode)

	dbConnection := server.CreateDBConnection()
	dataModels := data.NewModels(dbConnection)
	s := services.NewServices(dataModels)
	c := controllers.NewControllers(s)

	mainServer := server.CreateNew(ginEngine, dataModels, s, c)

	err := mainServer.Serve()
	if err != nil {
		fmt.Printf("Error during serving: %v\n", err)
	}
}
