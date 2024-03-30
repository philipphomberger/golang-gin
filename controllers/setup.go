package controllers

import (
	"ginapi/configs"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var validate = validator.New()

func GetCollection(name string) *mongo.Collection {
	return configs.GetCollection(configs.DB, name)
}
