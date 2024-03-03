package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName  string `gorm:"not null;" json:"firstName" `
	LastName   string `gorm:"not null;" json:"lastName" `
	Email      string `gorm:"not null;" json:"email" `
	Password   string `gorm:"not null;" json:"password" `
	RememberMe bool   `gorm:"not null;" json:"rememberMe" `
}

type Blog struct {
	gorm.Model
	Title     string `gorm:"not null;" json:"title" `
	Image_Url string `gorm:"not null;" json:"image_url"`
	Author    string `gorm:"not null; size:25" json:"author"`
	Body      string `gorm:"not null;" json:"body"`
}
