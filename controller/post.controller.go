package controller

import (
	"fmt"
	"net/http"

	"github.com/dorman99/go-gin-mongo/dto"
	"github.com/dorman99/go-gin-mongo/service"
	"github.com/gin-gonic/gin"
)

type PostController interface {
	Create(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type postController struct {
	postService service.PostService
}

func NewPostController(postService service.PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (p postController) Create(ctx *gin.Context) {
	createPostDto := dto.CreatePostDto{}
	err := ctx.ShouldBind(&createPostDto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	post, errI := p.postService.Create(createPostDto)
	if errI != nil {
		fmt.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": errI.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, post)

}

func (p postController) Find(ctx *gin.Context)    {}
func (p postController) FindAll(ctx *gin.Context) {}
