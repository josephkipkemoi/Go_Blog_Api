package main

import (
	"f1-blog/database"
	"f1-blog/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Load Environment Variables
func init() {
	env := os.Getenv("APP_ENV")
	if env == "development" || env == "" {
		err := godotenv.Load("./.env")
		if err != nil {
			log.Fatalf("error loading .env file: %s", err)
		}
	}
}

func main() {
	var port string = os.Getenv("APP_SERVER_PORT")

	database.ConnectDatabase() // connect to DB

	r := server.ConnectServer() // connect to server

	r.Run(port)
}
