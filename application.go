package main

import (
	"ginapi/configs"
	"ginapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create Gin Instance
	r := gin.Default()
	// Add Gin Routes
	routes.AlbumsRoute(r)
	routes.ArtistsRoute(r)
	routes.BandRoute(r)
	//run database
	configs.ConnectDB()
	// Run Server
	r.Run(":8080")
}
