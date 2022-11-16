package main

import (
	"github.com/My5z0n/FireDogCollector/Backend/cmd/api/server"
	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	backendRouter := server.NewBackendRouter(ginEngine)

	dbConnection := server.CreateDBConnection()
	mainServer := server.CreateNew(ginEngine, backendRouter)

	mainServer.Serve()
}
