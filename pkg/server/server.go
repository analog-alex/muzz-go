package server

import (
	"github.com/gin-gonic/gin"
	"muzz-service/pkg/controllers"
)

func Start() {
	// TODO fetch the port from the config
	port := ":8080"

	router := gin.Default()
	setupRoutes(router)

	if err := router.Run(port); err != nil {
		// no server running, crash the program with no survivors
		panic(err)
	}
}

// setupRoutes sets up the routes for the server
// centralized view over all the API endpoints exposed
func setupRoutes(router *gin.Engine) {
	// actuator-like endpoints
	router.GET("/health", controllers.Health)

	// user endpoints
	router.GET("/user", controllers.GetAll)
	router.POST("/user/create", controllers.Create)

	// auth endpoints
	router.POST("/login", controllers.Login)

	// match endpoints
	router.GET("/discover", controllers.Discover)
	router.POST("/swipe", controllers.Swipe)
}
