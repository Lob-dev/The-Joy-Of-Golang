package service

import (
	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/persistence"
)

type PostService interface {
	GetPosts() []persistence.Post
	AddPost(persistence.Post)
}

type postService struct {
	postRepository persistence.PostRepository
}

func Init(repository persistence.PostRepository) PostService {
	return &postService{
		postRepository: repository,
	}
}

func (this *postService) GetPosts() []persistence.Post {
	return this.postRepository.GetPosts()
}

func (this *postService) AddPost(post persistence.Post) {
	this.postRepository.AddPost(post)
}
