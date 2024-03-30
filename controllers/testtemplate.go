package controllers

import (
	"ginapi/models"
	"github.com/gin-gonic/gin"
)

func PostNewTest() gin.HandlerFunc {
	return CreateTemplate(GetCollection("template"), models.Song{}, models.Song{Title: "test", Artist: []string{"John Lennon"}})
}
