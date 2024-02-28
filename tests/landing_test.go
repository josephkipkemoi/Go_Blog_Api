package tests

import (
	"f1-blog/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestLandingApiPage(t *testing.T) {
	r := server.ConnectServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"F1 Secrets API v1"}`, w.Body.String())
}
