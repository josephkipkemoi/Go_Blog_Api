package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/handler"
	"f1-blog/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert"
)

func TestCannotRegisterNewUserWithInvalidPassword(t *testing.T) {
	r := server.ConnectServer()

	i := &handler.RegisterUserInput{
		FirstName:       "Joseph",
		LastName:        "Mwanzia",
		Email:           "jkemboe@gmail.com",
		Password:        "1235",
		ConfirmPassword: "12345",
		RememberMe:      true,
	}

	d, err := json.Marshal(i)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/user/register", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
