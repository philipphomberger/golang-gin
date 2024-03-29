package models

import (
	"testing"
)

func TestArtist(t *testing.T) {
	var artist = Artist{
		Name:    "Test",
		SurName: "Tim Ferris",
	}
	_ = artist
}
