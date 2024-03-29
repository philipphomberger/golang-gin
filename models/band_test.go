package models

import (
	"testing"
)

func TestBand(t *testing.T) {
	var members []string
	members = append(members, "Klaus")
	var band = Band{
		Name:    "Test",
		Members: members,
	}
	_ = band
}
