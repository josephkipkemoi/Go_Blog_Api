package database

import "fmt"

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

type UpdateCategoryInput struct {
	CategoryName string
}

func (c *Category) Update(id int, name string) error {
	res := DB.First(&c, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot update record")
	}

	DB.Model(&c).Update("categoryName", name)

	return nil
}

func (c *Category) Delete(id int) error {
	res := DB.Delete(&c, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot delete record")
	}

	return nil
}
