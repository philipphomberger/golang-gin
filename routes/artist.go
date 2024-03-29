package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func ArtistsRoute(router *gin.Engine) {
	router.POST("/artists", controllers.CreateAlbum())
	router.GET("/artists", controllers.GetAlbums())
	router.GET("/artists/:id", controllers.GetAlbum())
	router.DELETE("/artists/:id", controllers.DelAlbum())
	router.PUT("/artists/:id", controllers.PutAlbum())
}
