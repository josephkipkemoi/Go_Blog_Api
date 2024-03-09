package database

func Create(b BlogInterface) (*Blog, error) {
	blog, err := b.CreateBlog()
	return blog, err
}

func Index(b BlogInterface) ([]Blog, error) {
	blogs, err := b.GetBlogs()
	return blogs, err
}
