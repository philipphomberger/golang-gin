package controllers

import (
	"context"
	"fmt"
	"ginapi/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTemplate[T any](templateCollection *mongo.Collection, newModel func(T) T) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var model T

		//validate the request body
		if err := c.BindJSON(&model); err != nil {
			c.JSON(http.StatusBadRequest, responses.GenericResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&model); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.GenericResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		result, err := templateCollection.InsertOne(ctx, newModel(model))
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.GenericResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetTemplates(templateCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		result, err := templateCollection.Find(ctx, bson.D{{}})
		if err != nil {
			panic(err)
		}
		var results []bson.M
		if err := result.All(ctx, &results); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, responses.GenericResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": results}})
	}
}

func GetTemplate(templateCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result bson.M
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = templateCollection.FindOne(ctx, filter).Decode(&result)
		fmt.Println(result)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}
		c.JSON(http.StatusOK, responses.GenericResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func DelTemplate(templateCollection *mongo.Collection) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var result bson.M
		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = templateCollection.FindOneAndDelete(ctx, filter).Decode(&result)
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

func PutTemplate[T any](templateCollection *mongo.Collection, newModel func(T) T) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var model T
		var result T

		//validate the request body
		if err := c.BindJSON(&model); err != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&model); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AlbumResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		objectId, err := primitive.ObjectIDFromHex(c.Param("id"))
		filter := bson.D{{"_id", objectId}}
		err = templateCollection.FindOneAndReplace(ctx, filter, newModel(model)).Decode(&result)

		if err != nil {
			if err == mongo.ErrNoDocuments {
				return
			}
			panic(err)
		}

		c.JSON(http.StatusCreated, responses.AlbumResponse{Status: http.StatusCreated, Message: "success replaced", Data: map[string]interface{}{"data": result}})
	}
}
