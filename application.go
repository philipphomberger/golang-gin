package main

import (
	"fmt"
	"ginapi/configs"
	"ginapi/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	routes.AlbumsRoute(r)
	//run database
	var client = configs.ConnectDB()
	fmt.Println(configs.GetCollection(client, "albums"))
	r.Run(":8080")
}

func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to the API!",
	})
}

func getUser(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, " + name + "!",
	})
}
