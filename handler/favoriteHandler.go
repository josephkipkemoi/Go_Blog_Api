package handler

import (
	"f1-blog/database"
	"net/http"
	"strconv"

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
	f := &database.Favourite{}

	id, e := strconv.Atoi(ctx.Param("user_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	d, err := f.Index(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": d,
	})
}

func DeleteFavourite(ctx *gin.Context) {
	f := &database.Favourite{}

	user_id, e := strconv.Atoi(ctx.Param("user_id"))
	fav_id, _ := strconv.Atoi(ctx.Param("favourite_id"))

	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": e,
		})
		return
	}

	err := f.Delete(user_id, fav_id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "favorite deleted",
	})
}
