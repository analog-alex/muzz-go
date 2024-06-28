package main

import (
	"fmt"
	"log"
	"muzz-service/api"
	"muzz-service/config"
)

func main() {
	log.Println("Starting...")

	// get configs
	port := config.GetApplicationConfig().Port

	// get routes
	router := api.Routes()

	// start server
	log.Println("Booting up server on port", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		// no server running, crash the program with no survivors
		panic(err)
	}
}
