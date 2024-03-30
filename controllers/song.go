package controllers

import (
	"ginapi/models"
	"github.com/gin-gonic/gin"
)

func CreateSong() gin.HandlerFunc {
	return CreateTemplate[models.Song](GetCollection("songs"), func(song models.Song) models.Song {
		return song
	})
}

func GetSongs() gin.HandlerFunc {
	return GetTemplates(GetCollection("songs"))
}

func GetSong() gin.HandlerFunc {
	return GetTemplate(GetCollection("songs"))
}

func DelSong() gin.HandlerFunc {
	return DelTemplate(GetCollection("songs"))
}

func PutSong() gin.HandlerFunc {
	return PutTemplate[models.Song](GetCollection("songs"), func(song models.Song) models.Song {
		return song
	})
}
