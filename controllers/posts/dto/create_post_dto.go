package posts

type CreatePostDto struct {
	Title string `json:"title" binding:"required,max=255"`
}
