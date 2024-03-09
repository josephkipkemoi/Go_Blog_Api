package database

func (c *Category) Create() (*Category, error) {
	res := DB.Create(&c).Error
	if res != nil {
		return &Category{}, res
	}

	return c, nil
}
