package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func ArtistsRoute(router *gin.Engine) {
	router.POST("/artists", controllers.CreateArtist())
	router.GET("/artists", controllers.GetArtists())
	router.GET("/artists/:id", controllers.GetArtist())
	router.DELETE("/artists/:id", controllers.DelArtist())
	router.PUT("/artists/:id", controllers.PutArtist())
}
