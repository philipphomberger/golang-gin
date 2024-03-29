package models

type Song struct {
	Title  string   `json:"title,omitempty" validate:"required"`
	Artist []string `json:"artist,omitempty" validate:"required"`
	Album  string   `json:"album,omitempty" validate:"required"`
}
