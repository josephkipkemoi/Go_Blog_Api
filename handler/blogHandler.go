package handler

import (
	"f1-blog/database"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Index handler function returns all records from the database
func Index(ctx *gin.Context) {
	var frontEndUrl string = os.Getenv("FRONTEND_URL")

	ctx.Header("Access-Control-Allow-Origin", frontEndUrl)

	b := &database.Blog{}

	c_id, _ := strconv.Atoi(ctx.Query("c_id"))
	feat, _ := strconv.ParseBool(ctx.Query("featured"))

	d, f, c, err := b.GetBlogs(c_id, feat)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":     d,
		"featured": f,
		"category": c,
	})
}

// Create handler function creates a new record on the database with the provided fields
func Create(ctx *gin.Context) {
	b := &database.Blog{}

	validate = validator.New()
	e := validate.Struct(b)

	errs, ok := validationErrors(e)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": errs,
		})
		return
	}

	ctx.ShouldBindJSON(b)

	blog, e := database.Create(b)

	if e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating blog post",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": blog,
	})
}

// Show handler function returns a record that matches the provided Id, 404 error code is returned if record is not found
func Show(ctx *gin.Context) {
	b := &database.Blog{}

	id, e := strconv.Atoi(ctx.Param("blog_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "invalid blog id",
		})
		return
	}

	blog, err := b.GetBlogById(int(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": blog,
	})
}

// Delete handler function will delete a record from the database that matches given Id
func Delete(ctx *gin.Context) {
	b := &database.Blog{}

	id, e := strconv.Atoi(ctx.Param("blog_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "invalid blog id",
		})
		return
	}

	err := b.DeleteBlogById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"data": "record deleted from database",
	})
}

// Patch handler function will update a column or set of column in the database given the database record Id
func Patch(ctx *gin.Context) {
	b := &database.Blog{}
	p := &PatchBlog{}

	id, e := strconv.Atoi(ctx.Param("blog_id"))
	if e != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "invalid blog id",
		})
		return
	}

	ctx.ShouldBindJSON(&p)

	err := b.PatchBlogById(id, p)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{
		"data": b,
	})
}
