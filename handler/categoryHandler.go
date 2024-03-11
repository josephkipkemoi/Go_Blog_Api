package handler

import (
	"f1-blog/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func CreateCategory(ctx *gin.Context) {
	c := &database.Category{}

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

	ct, err := c.Create()
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"category": ct,
	})
}

func IndexCategory(ctx *gin.Context) {
	c := &database.Category{}

	d, err := c.Index()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": d,
	})
}
