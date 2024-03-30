package controllers

import (
	"ginapi/models"
	"github.com/gin-gonic/gin"
)

func CreateArtist() gin.HandlerFunc {
	return CreateTemplate[models.Artist](GetCollection("artists"), func(song models.Artist) models.Artist {
		return song
	})
}

func GetArtists() gin.HandlerFunc {
	return GetTemplates(GetCollection("artists"))
}

func GetArtist() gin.HandlerFunc {
	return GetTemplate(GetCollection("artists"))
}

func DelArtist() gin.HandlerFunc {
	return DelTemplate(GetCollection("artists"))
}

func PutArtist() gin.HandlerFunc {
	return PutTemplate[models.Artist](GetCollection("artists"), func(song models.Artist) models.Artist {
		return song
	})
}
