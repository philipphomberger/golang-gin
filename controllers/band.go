package controllers

import (
	"ginapi/models"
	"github.com/gin-gonic/gin"
)

func CreateBand() gin.HandlerFunc {
	return CreateTemplate[models.Band](GetCollection("bands"), func(song models.Band) models.Band {
		return song
	})
}

func GetBands() gin.HandlerFunc {
	return GetTemplates(GetCollection("bands"))
}

func GetBand() gin.HandlerFunc {
	return GetTemplate(GetCollection("bands"))
}

func DelBand() gin.HandlerFunc {
	return DelTemplate(GetCollection("bands"))
}

func PutBand() gin.HandlerFunc {
	return PutTemplate[models.Band](GetCollection("bands"), func(song models.Band) models.Band {
		return song
	})
}
