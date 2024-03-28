package controllers

import (
	"context"
	"ginapi/configs"
	"ginapi/models"
	"ginapi/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var albumsCollection *mongo.Collection = configs.GetCollection(configs.DB, "albums")
var validate = validator.New()

func CreateAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var album models.Album
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&album); err != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&album); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newAlbum := models.Album{
			ID:     primitive.NewObjectID(),
			Title:  album.Title,
			Artist: album.Artist,
			Price:  album.Price,
		}

		result, err := albumsCollection.InsertOne(ctx, newAlbum)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AlbumResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AlbumResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
