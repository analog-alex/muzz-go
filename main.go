package main

import (
	"fmt"
	"muzz-service/pkg/server"
)

func main() {
	server.Start()
	// boot up server
	fmt.Println("Booting up server in port 8080...")
}
