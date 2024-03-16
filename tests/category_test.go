package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/database"
	"f1-blog/handler"
	"f1-blog/server"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestCanGetCategories(t *testing.T) {
	r := server.ConnectServer()

	c := &database.Category{
		CategoryName: "Nascar Two ho",
	}

	c.Create()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/category", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCanUpdateCategoryById(t *testing.T) {
	r := server.ConnectServer()

	u := handler.UpdateCategoryInput{
		CategoryName: "Major League",
	}

	d, err := json.Marshal(u)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/api/v1/category/6", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestCanDeleteCategoryById(t *testing.T) {
	r := server.ConnectServer()

	c := &database.Category{
		CategoryName: "Some Name",
	}

	c.Create()

	c_id := strconv.Itoa(int(c.ID))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/category/"+c_id, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
