package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LandingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "F1 Secrets API v1",
	})
}
