package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectServer is responsible for setting up and launching the database
func ConnectDatabase() {
	var host string = os.Getenv("DB_HOST")
	var port string = os.Getenv("DB_PORT")
	var user string = os.Getenv("DB_USER")
	var password string = os.Getenv("DB_PASSWORD")
	var dbname string = os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s	password=%s	dbname=%s sslmode=disable", host, port, user, password, dbname)

	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error")
		log.Fatalf("Error Message: %s", err)
	}

	log.Println("********************************")
	log.Println("Database connected succesfully")
	log.Println("********************************")

	runMigrations(Db)

	DB = Db
}
