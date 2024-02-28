package server

import (
	"github.com/gin-gonic/gin"
)

// ConnectServer is responsible for setting up and launching the API server
func ConnectServer() *gin.Engine {

	r := gin.Default()

	Api(r) // Api routes

	return r
}
