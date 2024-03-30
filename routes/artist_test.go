package routes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetArtistsRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/artists", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetArtistRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/artists/6606e636ea20581b675f6cdf", nil)
	router.ServeHTTP(w, req)
	fmt.Println(w.Body.String())
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":200,"message":"success","data":{"data":{"_id":"6606e636ea20581b675f6cdf","id":"6606e636ea20581b675f6cde","name":"Lennon","surname":"John"}}}`, w.Body.String())
}
