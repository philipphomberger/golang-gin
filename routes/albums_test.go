package routes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAlbumsRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	fmt.Println(w.Body.String())
}

func TestGetAlbumRoute(t *testing.T) {
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/66053fd1e9647c9e6a7d8a25", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":200,"message":"success","data":{"data":{"_id":"66053fd1e9647c9e6a7d8a25","artist":"Beatles","id":"66053fd1e9647c9e6a7d8a24","price":12.99,"title":"Abbey Road"}}}`, w.Body.String())
}
