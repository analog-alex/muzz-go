package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"muzz-service/config"
	"muzz-service/pkg/handler"
	"muzz-service/pkg/middleware"
)

func Start() {
	port := config.GetApplicationConfig().Port

	router := gin.Default()
	setupRoutes(router)

	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		// no server running, crash the program with no survivors
		panic(err)
	}
}

// setupRoutes sets up the routes for the server
// centralized view over all the API endpoints exposed
func setupRoutes(router *gin.Engine) {
	// actuator-like endpoints
	router.GET("/health", handler.Health)

	// user endpoints
	router.GET("/user", middleware.AuthorizationMiddleware(), handler.GetAll)
	router.POST("/user/create", handler.Create)

	// auth endpoints
	router.POST("/login", handler.Login)

	// match endpoints
	router.GET("/discover", middleware.AuthorizationMiddleware(), handler.Discover)
	router.POST("/swipe", middleware.AuthorizationMiddleware(), handler.Swipe)
}
