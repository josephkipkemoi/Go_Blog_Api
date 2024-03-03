package handler

type RegisterUserInput struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
	Email           string `json:"email"`
	RememberMe      bool   `json:"rememberMe"`
}

type PatchBlog struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	Image_Url string `json:"image_url"`
	Body      string `json:"body"`
}
