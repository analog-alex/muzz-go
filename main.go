package main

import (
	"fmt"
	"log"
	"muzz-service/apis"
	"muzz-service/config"
)

func main() {
	log.Println("Booting up server in port 8080...")

	// get configs
	port := config.GetApplicationConfig().Port

	// get routes
	router := apis.Routes()

	// start server
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		// no server running, crash the program with no survivors
		panic(err)
	}
}
