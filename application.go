package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// Routes defined in the routes package
	routes := r.Group("/api")
	{
		routes.GET("/", getHome)
		routes.GET("/user/:name", getUser)
		routes.GET("/albums", getAlbums)
		routes.POST("/albums", postAlbum)
		routes.DELETE("/albums", deleteAlbum)
	}

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
