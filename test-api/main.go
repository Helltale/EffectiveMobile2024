package main

import (
	_ "github.com/helltale/EffectiveMobile2024/test-api/swagger/docs"

	"github.com/gin-gonic/gin"
	db "github.com/helltale/EffectiveMobile2024/test-api/db"
	"github.com/helltale/EffectiveMobile2024/test-api/routes"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Music API
// @version 1.0
// @description music API.
// @host localhost:8080
// @BasePath /
func main() {
	db.Connect()

	router := gin.Default()

	routes.RegisterRoutes(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8080")
}
