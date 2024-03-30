package routes

import (
	"ginapi/models"
	"github.com/mjarkk/mongomock"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAlbumsRoute(t *testing.T) {
	db := mongomock.NewDB()
	collection := db.Collection("Albums")
	err := collection.Insert(models.Album{
		Title:  "Let it Be",
		Artist: "The Beatles",
		Price:  12.99,
	})
	if err != nil {
		log.Fatal(err)
	}
	router := SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Equal(t, "pong", w.Body.String())
}
