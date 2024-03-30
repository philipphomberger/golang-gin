package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	AlbumsRoute(r)
	ArtistsRoute(r)
	BandRoute(r)
	SongsRoute(r)
	return r
}
