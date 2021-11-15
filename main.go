package main

import (
	"context"
	"fmt"

	"github.com/dorman99/go-gin-mongo/controller"
	"github.com/dorman99/go-gin-mongo/provider"
	"github.com/dorman99/go-gin-mongo/repository"
	"github.com/dorman99/go-gin-mongo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	client, database, err := provider.NewClient("mongodb://localhost:27017", "social")
	if err != nil {
		panic(err)
	}

	postRepository := repository.NewPostReposistory(database, "post")
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)

	defer client.Disconnect(context.Background())
	fmt.Println("---CONNECTED TO MONGO---")
	app := gin.Default()

	app.POST("/posts", postController.Create)
	app.Run()
}
