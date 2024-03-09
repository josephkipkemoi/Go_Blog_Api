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

func TestCanCreateRoles(t *testing.T) {
	r := server.ConnectServer()

	role := &database.Roles{
		RoleName: "Author",
	}

	d, err := json.Marshal(role)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/roles", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
