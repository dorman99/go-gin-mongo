package dto

type CreatePostDto struct {
	Title string `json:"title" binding:"required"`
}
