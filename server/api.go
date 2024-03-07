package server

import (
	"f1-blog/handler"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Api(r *gin.Engine) {
	var frontEndUrl string = os.Getenv("FRONTEND_URL")

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontEndUrl},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin, Content-Type, Token, Accept, X-Requested-With, withCredentials, Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Origin, Content-Type, Token, Accept, X-Requested-With, Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
