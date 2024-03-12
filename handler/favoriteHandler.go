package handler

import (
	"f1-blog/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateFavourite(ctx *gin.Context) {
	f := &database.Favourite{}

	validate := validator.New()
	e := validate.Struct(f)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.ShouldBindJSON(f)

	e = f.Create()
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": e,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "favourite added",
	})
}

func IndexFavourite(ctx *gin.Context) {
	panic("Not implemented")
}

func DeleteFavourite(ctx *gin.Context) {
	panic("Not implemented")
}
