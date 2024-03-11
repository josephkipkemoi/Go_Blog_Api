package handler

import (
	"f1-blog/database"
	"net/http"
	"strconv"

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

func IndexRoles(ctx *gin.Context) {
	r := &database.Roles{}

	d, err := r.Index()
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

func ShowRoles(ctx *gin.Context) {
	r := &database.Roles{}

	id, e := strconv.Atoi(ctx.Param("role_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	d, err := r.Show(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": d,
	})
}

func DeleteRoles(ctx *gin.Context) {
	r := &database.Roles{}

	id, e := strconv.Atoi(ctx.Param("role_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	err := r.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "role deleted",
	})
}
