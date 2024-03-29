package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Band struct {
	ID      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Members []string           `json:"ids,omitempty" validate:"required"`
}
