package routes

import (
	"ginapi/controllers"
	"github.com/gin-gonic/gin"
)

func BandRoute(router *gin.Engine) {
	router.POST("/bands", controllers.CreateBand())
	router.GET("/bands", controllers.GetBands())
	router.GET("/bands/:id", controllers.GetBand())
	router.DELETE("/bands/:id", controllers.DelBand())
	router.PUT("/bands/:id", controllers.PutBand())
}
