package posts

import (
	"context"
	"github.com/gin-gonic/gin"
	"inlove-app-server/controllers"
	posts "inlove-app-server/controllers/posts/dto"
	"inlove-app-server/db"
	prisma "inlove-app-server/prisma/db"
	"net/http"
)

type IPostController interface {
	controllers.IResourceController
}

type PostController struct {
	controllers.IResourceController
}

var (
	prismaClient = db.GetDB()
	contextB     = context.Background()
)

func (pc PostController) Create(c *gin.Context) {
	var payload posts.CreatePostDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := prismaClient.Post.CreateOne(
		prisma.Post.Title.Set(payload.Title),
		prisma.Post.Published.Set(false),
	).Exec(contextB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc PostController) List(c *gin.Context) {
	postResults, err := prismaClient.Post.FindMany().Exec(contextB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postResults)
}

func (pc PostController) RootName() string {
	return "post"
}
