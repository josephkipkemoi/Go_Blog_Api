package database

import "fmt"

type BlogInterface interface {
	CreateBlog() (*Blog, error)
	GetBlogs() ([]Blog, error)
	GetBlogById(int) (*Blog, error)
	DeleteBlogById(int) error
	PatchBlogById(int, interface{}) error
}

func (b *Blog) CreateBlog() (*Blog, error) {
	err := DB.Create(&b).Error

	if err != nil {
		return &Blog{}, err
	}

	return b, nil
}

func (b *Blog) GetBlogs() ([]Blog, error) {
	var blog []Blog
	res := DB.Order("id desc").Limit(20).Find(&blog)

	if res.Error != nil {
		return blog, res.Error
	}

	return blog, nil
}

func (b *Blog) GetBlogById(id int) (*Blog, error) {
	res := DB.Find(&b, id)

	if res.Error != nil || res.RowsAffected == 0 {
		return &Blog{}, fmt.Errorf("%s", "record not found")
	}

	return b, nil
}

func (b *Blog) DeleteBlogById(id int) error {
	res := DB.Delete(&b, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot delete record")
	}

	return nil
}

func (b *Blog) PatchBlogById(id int, bg interface{}) error {
	res := DB.First(&b, id)
	if res.Error != nil || res.RowsAffected == 0 {
		return fmt.Errorf("%s", "not found: cannot update record")
	}

	DB.Model(&b).Updates(bg)

	return nil
}
