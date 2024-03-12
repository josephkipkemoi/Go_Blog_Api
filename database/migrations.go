package database

import (
	"log"

	"gorm.io/gorm"
)

func runMigrations(Db *gorm.DB) {
	log.Println("********************************")
	log.Println("Running migrations")
	log.Println("********************************")

	Db.AutoMigrate(
		&Favourite{},
		&Contact{},
		&Roles{},
		&User{},
		&Blog{},
		&Category{},
	)

	log.Println("********************************")
	log.Println("Migrations complete")
	log.Println("********************************")
}
