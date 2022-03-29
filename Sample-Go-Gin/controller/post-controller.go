package controller

import (
	"fmt"

	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/persistence"
	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	GetPosts() []persistence.Post
	AddPost(context *gin.Context) string
}

type postController struct {
	postService service.PostService
}

func Init(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (this *postController) GetPosts() []persistence.Post {
	return this.postService.GetPosts()
}

func (this *postController) AddPost(context *gin.Context) string {
	var post persistence.Post
	context.BindJSON(&post)
	this.postService.AddPost(post)
	return fmt.Sprintf("Add Post! title = %s, description = %s", post.Title, post.Description)
}
