package database

import "fmt"

func (u *User) RegisterUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, nil
	}
	return u, nil
}

func (u *User) LoginUser(email string, password string) (*User, error) {
	DB.First(&u, "email = ?", email)
	if u.Password != password {
		return &User{}, fmt.Errorf("%s", "email or password do not match")
	}
	return u, nil
}
