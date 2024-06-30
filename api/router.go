package api

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/handlers"
	"muzz-service/pkg/middleware"
)

func Routes() *gin.Engine {
	router := gin.Default()

	// actuator-like endpoints
	router.GET("/health", handlers.Health)

	// user endpoints
	router.POST("/user/create", handlers.CreateUser)

	// auth endpoints
	router.POST("/login", handlers.Login)

	// match endpoints
	router.GET("/discover", middleware.AuthorizationMiddleware(), handlers.Discover)
	router.POST("/swipe", middleware.AuthorizationMiddleware(), handlers.Swipe)

	return router
}
