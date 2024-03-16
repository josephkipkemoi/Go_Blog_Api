package handler

import (
	"f1-blog/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type UpdateCategoryInput struct {
	CategoryName string
}

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

func UpdateCategory(ctx *gin.Context) {
	i := &UpdateCategoryInput{}
	c := &database.Category{}
	c_id, _ := strconv.Atoi(ctx.Param("category_id"))

	validate := validator.New()
	e := validate.Struct(c)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	ctx.ShouldBindJSON(i)

	err := c.Update(c_id, i.CategoryName)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "category name updated",
	})

}
