package main

import (
	"f1-blog/database"
	"f1-blog/server"
)

const (
	port string = ":5050"
)

func main() {
	database.ConnectDatabase() // connect to DB

	r := server.ConnectServer() // connect to server

	r.Run(port)
}
