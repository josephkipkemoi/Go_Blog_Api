package database

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	RoleName string `gorm:"unique;not null;" json:"roleName"`
}

type User struct {
	gorm.Model
	RoleId     int    `gorm:"not null;default:3" json:"roleId"`
	FirstName  string `gorm:"not null;" json:"firstName"`
	LastName   string `gorm:"not null;" json:"lastName"`
	Email      string `gorm:"unique" json:"email"`
	Password   string `gorm:"not null;" json:"-"`
	RememberMe bool   `gorm:"not null;" json:"rememberMe"`
}

type Blog struct {
	gorm.Model
	CategoryId int    `json:"categoryId" `
	Title      string `gorm:"not null;" json:"title" `
	Featured   bool   `json:"featured"`
	Image_Url  string `gorm:"not null;" json:"image_url"`
	Author     string `gorm:"not null; size:25" json:"author"`
	Body       string `gorm:"not null;" json:"body"`
}

type Category struct {
	gorm.Model
	CategoryName string `gorm:"unique;not null;" json:"categoryName"`
}

type Contact struct {
	gorm.Model
	FirstName string `gorm:"not null;" json:"firstName"`
	LastName  string `gorm:"not null;" json:"lastName"`
	Email     string `gorm:"not null;" json:"email"`
	Message   string `gorm:"not null;" json:"message"`
}

type Favourite struct {
	gorm.Model
	BlogId int `gorm:"not null;" json:"blogId"`
	UserId int `gorm:"not null;" json:"userId"`
}
