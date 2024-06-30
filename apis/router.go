package apis

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/handler"
	"muzz-service/pkg/middleware"
)

func Routes() *gin.Engine {
	router := gin.Default()

	// actuator-like endpoints
	router.GET("/health", handler.Health)

	// user endpoints
	router.POST("/user/create", handler.CreateUser)

	// auth endpoints
	router.POST("/login", handler.Login)

	// match endpoints
	router.GET("/discover", middleware.AuthorizationMiddleware(), handler.Discover)
	router.POST("/swipe", middleware.AuthorizationMiddleware(), handler.Swipe)

	return router
}
