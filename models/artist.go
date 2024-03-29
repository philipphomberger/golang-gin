package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Artist struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	SurName string             `json:"surname,omitempty" validate:"required"`
}
