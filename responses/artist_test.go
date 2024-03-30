package responses

import (
	"testing"
)

func TestArtistResponse(t *testing.T) {
	var data map[string]interface{}
	var response = ArtistResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
	_ = response
}
