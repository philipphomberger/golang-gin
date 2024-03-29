package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func SongsRoute(router *gin.Engine) {
	router.POST("/songs", controllers.CreateSong())
	router.GET("/songs", controllers.GetSongs())
	router.GET("/songs/:id", controllers.GetSong())
	router.DELETE("/songs/:id", controllers.DelSong())
	router.PUT("/songs/:id", controllers.PutSong())
}
