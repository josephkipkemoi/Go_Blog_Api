package database

func (c *Category) Create() (*Category, error) {
	res := DB.Create(c).Error
	if res != nil {
		return &Category{}, res
	}

	return c, nil
}

func (c *Category) Index() ([]Category, error) {
	var category []Category

	res := DB.Order("id").Find(&category)

	if res.Error != nil {
		return category, res.Error
	}

	return category, nil
}
