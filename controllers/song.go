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

var songsCollection *mongo.Collection = configs.GetCollection(configs.DB, "artists")

func CreateSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var songs models.Song
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&songs); err != nil {
			c.JSON(http.StatusBadRequest, responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&songs); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSong := models.Song{
			Title:  songs.Title,
			Artist: songs.Artist,
			Album:  songs.Album,
		}

		result, err := artistsCollection.InsertOne(ctx, newSong)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.SongResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.SongResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetSongs() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := songsCollection.Find(ctx, bson.D{{}})
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
		c.JSON(http.StatusOK, responses.SongResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": results}})
	}
}

func GetSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		fmt.Println(c.Param("id"))
		var result models.Song
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = artistsCollection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.SongResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DelSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Artist
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = songsCollection.FindOneAndDelete(ctx, filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.SongResponse{Status: http.StatusOK, Message: "success delete", Data: map[string]interface{}{"data": result}})
	}
}

func PutSong() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Song
		var song = models.Song{}

		//validate the request body
		if err := c.BindJSON(&song); err != nil {
			c.JSON(http.StatusBadRequest, responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&song); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.SongResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newSong := models.Song{
			Title:  song.Title,
			Artist: song.Artist,
			Album:  song.Album,
		}

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = albumsCollection.FindOneAndReplace(ctx, filter, newSong).Decode(&result)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}

		c.JSON(http.StatusCreated, responses.SongResponse{Status: http.StatusCreated, Message: "success replaced", Data: map[string]interface{}{"data": newSong}})
	}
}
