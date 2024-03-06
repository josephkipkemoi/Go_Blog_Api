package handler

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

// createJWTToken function creates a JWT Token and provides the client for future request validations
func CreateJWTToken(username string) (string, error) {
	var (
		key []byte
		t   *jwt.Token
	)

	key = []byte("f1Secrets")           // load from .env
	t = jwt.New(jwt.SigningMethodHS256) // create new token

	claims := t.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["authorized"] = true
	claims["username"] = username

	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return s, nil
}

// VerifyToken function is responsible for checking whether the given token by the client is valid
func VerifyToken(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return tokenString, nil
	})
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return false, err
	}

	c := token.Claims.(jwt.MapClaims)
	authorized := c["authorized"]

	return authorized, nil
}

// validationErrors returns found errors stored in a slice and true if errors are found empty slice and false otherwise
func validationErrors(err error) ([]string, bool) {
	errs := []string{}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errs, false
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println()
			fmt.Println(err)
			errs = append(errs, err.Field()+" Field is required")
		}
		return errs, false
	}
	return errs, true
}

// comparePasswords function confirm if the passwords the user provides when registering are equal
func comparePasswords(str1, str2 string) bool {
	if str1 == str2 {
		return true
	} else {
		return false
	}
}
