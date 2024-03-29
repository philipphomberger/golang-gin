package controllers

import (
	"context"
	"fmt"
	"ginapi/configs"
	"ginapi/models"
	"ginapi/responses"
	"go.mongodb.org/mongo-driver/bson"
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

func GetAlbums() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := albumsCollection.Find(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}
		var results []bson.M
		if err := result.All(ctx, &results); err != nil {
			panic(err)
		}
		for _, doc := range results {
			fmt.Println(doc)
		}
		c.JSON(http.StatusOK, responses.AlbumResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": results}})
	}
}

func GetAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println(c.Param("id"))
		var result models.Album
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = albumsCollection.FindOne(ctx, filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.AlbumResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DelAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Album
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = albumsCollection.FindOneAndDelete(ctx, filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.AlbumResponse{Status: http.StatusOK, Message: "success delete", Data: map[string]interface{}{"data": result}})
	}
}

func PutAlbum() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Album
		var album = models.Album{}

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

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = albumsCollection.FindOneAndReplace(ctx, filter, newAlbum).Decode(&result)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}

		c.JSON(http.StatusCreated, responses.AlbumResponse{Status: http.StatusCreated, Message: "success replaced", Data: map[string]interface{}{"data": newAlbum}})
	}
}
