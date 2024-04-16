package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.Default()

// Starts the server
func Run() {
	addRoutes()
	router.Run(":8080")
}

// Adds all the routes for the application
func addRoutes() {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("api")
	v1 := api.Group("/v1")
	AddTodoRoutes(v1)
}
