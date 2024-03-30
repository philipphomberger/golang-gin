package responses

import (
	"testing"
)

func TestAlbumResponse(t *testing.T) {
	var data map[string]interface{}
	var response = AlbumResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
	_ = response
}
