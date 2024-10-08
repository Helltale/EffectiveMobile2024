package main

import (
	_ "github.com/helltale/EffectiveMobile2024/test-api/swagger/docs"

	"github.com/gin-gonic/gin"
	db "github.com/helltale/EffectiveMobile2024/test-api/internal/db"
	"github.com/helltale/EffectiveMobile2024/test-api/internal/logger"
	"github.com/helltale/EffectiveMobile2024/test-api/internal/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Music API
// @version 1.0
// @description music API.
// @host localhost:8080
// @BasePath /
func main() {
	logger.Init()

	db.Connect()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	logger.Log.Info("star server on port 8080")
	router.Run(":8080")
}
