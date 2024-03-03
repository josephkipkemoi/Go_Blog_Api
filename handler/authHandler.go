package handler

import (
	"f1-blog/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func AuthRegister(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "http://127.0.0.1:3000")

	i := &RegisterUserInput{}
	usr := &database.User{}

	validate = validator.New()
	e := validate.Struct(i)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	ctx.ShouldBindJSON(i)

	passwordVerified := comparePasswords(i.Password, i.ConfirmPassword)
	if !passwordVerified {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "passwords do not match",
		})
		return
	}

	usr.FirstName = i.FirstName
	usr.LastName = i.LastName
	usr.Password = i.Password
	usr.Email = i.Email
	usr.RememberMe = i.RememberMe

	u, e := usr.RegisterUser()
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot register user",
		})
		return
	}

	tokenString, err := createJWTToken(usr.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "authorization failure",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"user":  u,
		"token": tokenString,
	})
}
