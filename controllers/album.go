package controllers

import (
	"ginapi/models"
	"github.com/gin-gonic/gin"
)

func CreateAlbum() gin.HandlerFunc {
	return CreateTemplate[models.Album](GetCollection("albums"), func(song models.Album) models.Album {
		return song
	})
}

func GetAlbums() gin.HandlerFunc {
	return GetTemplates(GetCollection("albums"))
}

func GetAlbum() gin.HandlerFunc {
	return GetTemplate(GetCollection("albums"))
}

func DelAlbum() gin.HandlerFunc {
	return DelTemplate(GetCollection("albums"))
}

func PutAlbum() gin.HandlerFunc {
	return PutTemplate[models.Album](GetCollection("albums"), func(song models.Album) models.Album {
		return song
	})
}
