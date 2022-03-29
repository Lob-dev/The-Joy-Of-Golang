package main

import (
	"net/http"

	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/controller"
	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/persistence"
	"github.com/Lob-dev/The-Joy-Of-Go/Sample-Go-Gin/service"
	"github.com/gin-gonic/gin"
)

var (
	postRepository persistence.PostRepository = persistence.Init()
	postService    service.PostService        = service.Init(postRepository)
	postController controller.PostController  = controller.Init(postService)
)

func main() {
	router := gin.Default()
	router.GET("/api/status", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"Message": "running",
		})
	})

	router.GET("/api/posts", func(context *gin.Context) {
		context.JSON(http.StatusOK, postController.GetPosts())
	})

	router.POST("/api/posts", func(context *gin.Context) {
		context.JSON(http.StatusCreated, postController.AddPost(context))
	})
	router.Run(":8080")
}
