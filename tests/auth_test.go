package tests

import (
	"bytes"
	"encoding/json"
	"f1-blog/database"
	"f1-blog/handler"
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

func TestCanRegisterNewUserWithValidCredentials(t *testing.T) {
	r := server.ConnectServer()

	i := &handler.RegisterUserInput{
		FirstName:       "Joseph",
		LastName:        "Mwanzia",
		Email:           "jkemboe@gmail.com",
		Password:        "12345",
		ConfirmPassword: "12345",
		RememberMe:      true,
	}

	d, err := json.Marshal(i)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/user/register", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestRegisteredUserCannotLogInWithInvalidCredentials(t *testing.T) {
	r := server.ConnectServer()

	u := &database.User{
		FirstName:  "joseph",
		LastName:   "ngetich",
		Email:      "jkemboe@gmail.com",
		Password:   "12345",
		RememberMe: true,
	}

	u.RegisterUser()

	i := &handler.LoginUserInput{
		Email:      "jkemboe@gmail.com",
		Password:   "123456",
		RememberMe: true,
	}

	d, err := json.Marshal(i)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/user/login", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestRegisteredUserCanLogInWithValidCredentials(t *testing.T) {
	r := server.ConnectServer()

	u := &database.User{
		FirstName:  "joseph",
		LastName:   "ngetich",
		Email:      "jkemboe@gmail.com",
		Password:   "12345",
		RememberMe: true,
	}

	u.RegisterUser()

	i := &handler.LoginUserInput{
		Email:      "jkemboe@gmail.com",
		Password:   "12345",
		RememberMe: true,
	}

	d, err := json.Marshal(i)
	checkErr(err)

	body := bytes.NewReader(d)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/auth/user/login", body)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCanAuthenticateUser(t *testing.T) {
	r := server.ConnectServer()

	u := &database.User{
		FirstName:  "joseph",
		LastName:   "ngetich",
		Email:      "jkemboe@gmail.com",
		Password:   "12345",
		RememberMe: true,
	}

	u.RegisterUser()

	tokenString, _ := handler.CreateJWTToken(u.Email)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/auth/user/verify", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		"Token":        {tokenString},
	}
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
