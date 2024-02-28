package handler

type PatchBlog struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Image_Url string `json:"image_url"`
	Body      string `json:"body"`
}
