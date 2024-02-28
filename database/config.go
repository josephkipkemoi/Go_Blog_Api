package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "jkemboe"
	password = "commandme007!@~"
	dbname   = "f1secretsdb"
)

// ConnectServer is responsible for setting up and launching the database
func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s	password=%s	dbname=%s sslmode=disable", host, port, user, password, dbname)

	Db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		fmt.Println("Database connection error")
		log.Fatalf("Error Message: %s", err)
	}

	fmt.Println("********************************")
	fmt.Println("Database connected succesfully")
	fmt.Println("********************************")

	Db.AutoMigrate(
		&Blog{},
	)

	DB = Db
}
