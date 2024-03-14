package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/database"
	"f1-blog/server"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/go-playground/assert"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading .env file: ", err)
	}

	database.ConnectDatabase()
}

/*
CRUD TESTS FOR BLOG API
This test suite is responsible for checking the backend logic for:
1. Creating resources on the database
2. Reading resources from the database
3. Updating a resource from the database
4. Deleting a resource from the database
*/
func TestAdminCanPostBlog(t *testing.T) {
	r := server.ConnectServer()

	b := &database.Blog{
		Title:      "Test Title",
		CategoryId: 1,
		Featured:   true,
		Author:     "Joseph Maasai",
		Image_Url:  "https://image.url",
		Body:       "Some blog body test",
	}

	d, err := json.Marshal(b)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/blog/new", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCanGetBlogPosts(t *testing.T) {
	r := server.ConnectServer()

	b := &database.Blog{
		Title:      "Test Title",
		CategoryId: 1,
		Featured:   true,
		Author:     "Joseph Maasai",
		Image_Url:  "https://image.url",
		Body:       "Some blog body test",
	}

	b.CreateBlog()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/blog?c_id=1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCanGetBlogPostById(t *testing.T) {
	r := server.ConnectServer()

	b := &database.Blog{
		Title:     "Test Title",
		Author:    "Joseph Maasai",
		Image_Url: "https://image.url",
		Body:      "Some blog body test",
	}

	b.CreateBlog()

	blogId := strconv.Itoa(int(b.ID))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/blogs/"+blogId, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAdminCanDeleteBlogById(t *testing.T) {
	r := server.ConnectServer()

	b := &database.Blog{
		Title:     "Test Title",
		Author:    "Joseph Maasai",
		Image_Url: "https://image.url",
		Body:      "Some blog body test",
	}

	b.CreateBlog()

	blogId := strconv.Itoa(int(b.ID))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/blogs/"+blogId+"/delete", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestAdminCanUpdateBlogById(t *testing.T) {
	r := server.ConnectServer()

	b := &database.Blog{
		Title:     "Test Title",
		Author:    "Joseph Maasai",
		Image_Url: "https://image.url",
		Body:      "Some blog body test",
	}

	b.CreateBlog()

	blogId := strconv.Itoa(int(b.ID))

	d, err := json.Marshal(&database.Blog{
		Title: "Updated Title",
	})
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PATCH", "/api/v1/blogs/"+blogId+"/patch", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
