package database

func (u *User) RegisterUser() (*User, error) {
	err := DB.Create(&u).Error

	if err != nil {
		return &User{}, nil
	}
	return u, nil
}
