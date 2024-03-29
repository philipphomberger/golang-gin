package models

import (
	"testing"
)

func TestSong(t *testing.T) {
	var artists []string
	var song = Song{
		Title:  "Test",
		Artist: append(artists, "Klaus"),
		Album:  "Test",
	}
	_ = song
}
