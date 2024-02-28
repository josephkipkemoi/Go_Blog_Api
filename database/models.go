package database

import "gorm.io/gorm"

type Blog struct {
	gorm.Model
	Title     string `gorm:"not null; size:50" json:"title" `
	Image_Url string `gorm:"not null; size:255" json:"image_url"`
	Author    string `gorm:"not null; size:25" json:"author"`
	Body      string `gorm:"not null;" json:"body"`
}
