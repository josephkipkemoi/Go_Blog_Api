package handler

import (
	"f1-blog/database"
	"net/http"
	"strconv"

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

func IndexContact(ctx *gin.Context) {
	c := &database.Contact{}

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

func ShowContact(ctx *gin.Context) {
	c := &database.Contact{}

	id, e := strconv.Atoi(ctx.Param("contact_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	d, err := c.Show(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": d,
	})
}
