package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func AlbumsRoute(router *gin.Engine) {
	router.POST("/albums", controllers.CreateAlbum())
	router.GET("/albums", controllers.GetAlbums())
	router.GET("/albums/:id", controllers.GetAlbum())
	router.DELETE("/albums/:id", controllers.DelAlbum())
	router.PUT("/albums/:id", controllers.PutAlbum())
}
