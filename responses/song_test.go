package responses

import (
	"testing"
)

func TestSongResponse(t *testing.T) {
	var data map[string]interface{}
	var response = SongResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
	_ = response
}
