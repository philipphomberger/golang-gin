package main

import (
	"ginapi/configs"
	"ginapi/routes"
	"log"
)

func main() {
	// Create Gin Instance
	r := routes.SetupRouter()
	//run database
	configs.ConnectDB()
	// Run Server
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
