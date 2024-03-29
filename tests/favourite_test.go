package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/database"
	"f1-blog/server"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading .env file: ", err)
	}
}

func init() {
	database.ConnectDatabase()
}

func TestClientCanPostFavourites(t *testing.T) {
	s := server.ConnectServer()

	fav := &database.Favourite{
		BlogId: 1,
		UserId: 2,
	}

	d, err := json.Marshal(fav)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/favourites", body)
	s.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestClientCanGetFavourites(t *testing.T) {
	r := server.ConnectServer()

	fav := &database.Favourite{
		UserId: 1,
		BlogId: 23,
	}

	fav.Create()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/favourites/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestClientCanDeleteFavouritesById(t *testing.T) {
	r := server.ConnectServer()

	fav := &database.Favourite{
		UserId: 1,
		BlogId: 23,
	}

	fav.Create()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/favourites/2/users/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
