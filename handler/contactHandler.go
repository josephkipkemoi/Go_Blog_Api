package handler

import (
	"f1-blog/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateContact(ctx *gin.Context) {
	c := &database.Contact{}

	validate := validator.New()
	e := validate.Struct(c)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.ShouldBindJSON(c)

	e = c.Create()
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": e,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "message sent",
	})
}
