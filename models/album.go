package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Album struct {
	ID     primitive.ObjectID `json:"id,omitempty"`
	Title  string             `json:"title,omitempty" validate:"required"`
	Artist string             `json:"artist,omitempty" validate:"required"`
	Price  float64            `json:"price,omitempty" validate:"required"`
}
