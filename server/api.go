package server

import (
	"f1-blog/handler"

	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	// Landing Route
	r.GET("/", handler.LandingHandler)

	// User Routes
	r.POST("/api/v1/auth/user/register", handler.AuthRegister)
	r.POST("/api/v1/auth/user/login", handler.AuthLogin)
	r.GET("/api/v1/auth/user/verify", handler.AuthVerify)

	// Blog Routes
	r.POST("/api/v1/blog/new", handler.Create)
	r.GET("/api/v1/blog", handler.Index)
	r.GET("/api/v1/blogs/:blog_id", handler.Show)
	r.DELETE("/api/v1/blogs/:blog_id/delete", handler.Delete)
	r.PATCH("/api/v1/blogs/:blog_id/patch", handler.Patch)
}
