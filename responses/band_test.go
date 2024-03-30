package responses

import (
	"testing"
)

func TestBandResponse(t *testing.T) {
	var data map[string]interface{}
	var response = BandResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	}
	_ = response
}
