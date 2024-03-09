package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/database"
	"f1-blog/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestCanCreateCategory(t *testing.T) {
	r := server.ConnectServer()

	c := &database.Category{
		CategoryName: "Nascar",
	}

	d, err := json.Marshal(c)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/category", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
