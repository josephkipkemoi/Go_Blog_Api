package database

import (
	"fmt"
)

type BlogInterface interface {
	CreateBlog() (*Blog, error)
	GetBlogs(int, bool) ([]Blog, []Blog, string, error)
	GetBlogById(int) (*Blog, error)
	DeleteBlogById(int) error
	PatchBlogById(int, interface{}) error
}

func (b *Blog) CreateBlog() (*Blog, error) {
	err := DB.Create(&b).Error

	DB.Table("blogs").Take(b)
	if err != nil {
		return &Blog{}, err
	}

	return b, nil
}

func (b *Blog) GetBlogs(c_id int, f bool) ([]Blog, []Blog, string, error) {
	c := Category{}
	var blog []Blog
	var featured []Blog
	DB.Table("categories").Find(&c, c_id)
	res := DB.Where("category_id = ?", c_id).Order("id desc").Limit(5).Find(&blog)
	DB.Where("featured = ?", f).Where("category_id = ?", c_id).Order("id desc").Limit(3).Find(&featured)

	if res.Error != nil {
		return blog, featured, "", res.Error
	}

	return blog, featured, c.CategoryName, nil
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
