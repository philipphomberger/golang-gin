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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var artistsCollection *mongo.Collection = configs.GetCollection(configs.DB, "artists")

func CreateArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var artist models.Artist
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&artist); err != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&artist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newArtist := models.Artist{
			ID:      primitive.NewObjectID(),
			Name:    artist.Name,
			SurName: artist.SurName,
		}

		result, err := artistsCollection.InsertOne(ctx, newArtist)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AlbumResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AlbumResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetArtists() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := artistsCollection.Find(ctx, bson.D{{}})
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

func GetArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println(c.Param("id"))
		var result models.Album
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = artistsCollection.FindOne(ctx, filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.ArtistResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DelArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Artist
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = artistsCollection.FindOneAndDelete(ctx, filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.ArtistResponse{Status: http.StatusOK, Message: "success delete", Data: map[string]interface{}{"data": result}})
	}
}

func PutArtist() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Artist
		var artist = models.Artist{}

		//validate the request body
		if err := c.BindJSON(&artist); err != nil {
			c.JSON(http.StatusBadRequest, responses.ArtistResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&artist); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ArtistResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newArtist := models.Artist{
			ID:      primitive.NewObjectID(),
			Name:    artist.Name,
			SurName: artist.SurName,
		}

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = albumsCollection.FindOneAndReplace(ctx, filter, newArtist).Decode(&result)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}

		c.JSON(http.StatusCreated, responses.ArtistResponse{Status: http.StatusCreated, Message: "success replaced", Data: map[string]interface{}{"data": newArtist}})
	}
}
