package service

import (
	"github.com/dorman99/go-gin-mongo/dto"
	"github.com/dorman99/go-gin-mongo/entity"
	"github.com/dorman99/go-gin-mongo/repository"
)

type PostService interface {
	Create(createPostDto dto.CreatePostDto) (entity.Post, error)
}

type postService struct {
	postRepository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &postService{
		postRepository: postRepository,
	}
}

func (p *postService) Create(createPostDto dto.CreatePostDto) (entity.Post, error) {
	return p.postRepository.Create(createPostDto)
}
