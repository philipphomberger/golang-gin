package models

import (
	"testing"
)

func TestAlbum(t *testing.T) {
	var album = Album{
		Title:  "Test",
		Artist: "Tim Ferris",
		Price:  12.99,
	}
	_ = album
}
