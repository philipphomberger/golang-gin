package main

import (
	"ginapi/configs"
	"ginapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.AlbumsRoute(r)
	//run database
	configs.ConnectDB()
	r.Run(":8080")
}
