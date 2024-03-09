package handler

import (
	"f1-blog/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func CreateRole(ctx *gin.Context) {
	r := &database.Roles{}

	validate := validator.New()
	e := validate.Struct(r)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.ShouldBindJSON(r)

	role, e := r.Create()
	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": e,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"role": role,
	})
}
