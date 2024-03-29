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

var bandCollection *mongo.Collection = configs.GetCollection(configs.DB, "bands")

func CreateBand() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var band models.Band
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&band); err != nil {
			c.JSON(http.StatusBadRequest, responses.BandResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&band); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BandResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newBand := models.Band{
			ID:      primitive.NewObjectID(),
			Name:    band.Name,
			Members: band.Members,
		}

		result, err := artistsCollection.InsertOne(ctx, newBand)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BandResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BandResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetBands() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := bandCollection.Find(ctx, bson.D{{}})
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
		c.JSON(http.StatusOK, responses.BandResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": results}})
	}
}

func GetBand() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Band
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = artistsCollection.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.BandResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DelBand() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Band
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = artistsCollection.FindOneAndDelete(ctx, filter).Decode(&result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.BandResponse{Status: http.StatusOK, Message: "success delete", Data: map[string]interface{}{"data": result}})
	}
}

func PutBand() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result models.Band
		var band = models.Band{}

		//validate the request body
		if err := c.BindJSON(&band); err != nil {
			c.JSON(http.StatusBadRequest, responses.ArtistResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&band); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.ArtistResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newArtist := models.Band{
			ID:      primitive.NewObjectID(),
			Name:    band.Name,
			Members: band.Members,
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

		c.JSON(http.StatusCreated, responses.BandResponse{Status: http.StatusCreated, Message: "success replaced", Data: map[string]interface{}{"data": newArtist}})
	}
}
