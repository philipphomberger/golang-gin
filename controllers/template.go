package controllers

import (
	"context"
	"ginapi/configs"
	"ginapi/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollection(name string) *mongo.Collection {
	return configs.GetCollection(configs.DB, name)
}

func CreateTemplate[T any](templateCollection *mongo.Collection, model T, newModel T) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

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

		result, err := templateCollection.InsertOne(ctx, newModel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.GenericResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.GenericResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}
