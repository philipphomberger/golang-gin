package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func AlbumsRoute(router *gin.Engine) {
	router.POST("/albums", controllers.CreateAlbum())
}
