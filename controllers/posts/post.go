package posts

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"inlove-app-server/constants/keys"
	"inlove-app-server/controllers"
	posts "inlove-app-server/controllers/posts/dto"
	prisma "inlove-app-server/prisma/db"
	"inlove-app-server/types"
	"net/http"
)

type IPostController interface {
	controllers.IResourceController
}

type PostController struct {
	controllers.IResourceController
	prisma *prisma.PrismaClient
}

var (
	contextB = context.Background()
)

// NewPostController returns a new instance of PostController.
func NewPostController(prisma *prisma.PrismaClient) IPostController {
	return &PostController{
		prisma: prisma,
	}
}

type LogInfo struct {
	Service string `json:"service"`
}

func (pc PostController) Create(c *gin.Context) {
	var payload posts.CreatePostDto
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, _ := c.Get(keys.UserKey)
	user := userData.(types.User)

	storedUser, err := pc.prisma.User.FindFirst(
		prisma.User.SupabaseUserID.Equals(user.Id),
	).Exec(contextB)

	if storedUser == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	logInfo := &LogInfo{
		Service: "deployment/api",
	}
	infoBytes, err := json.Marshal(logInfo)
	if err != nil {
		panic(err)
	}

	post, err := pc.prisma.Post.CreateOne(
		prisma.Post.Title.Set(payload.Title),
		prisma.Post.Published.Set(false),
		prisma.Post.IsPublic.Set(false),
		prisma.Post.User.Link(
			prisma.User.ID.Equals(storedUser.ID),
		),
		prisma.Post.Data.Set(prisma.JSON(infoBytes)),
	).Exec(contextB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (pc PostController) List(c *gin.Context) {
	postResults, err := pc.prisma.Post.FindMany().Exec(contextB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, postResults)
}

func (pc PostController) RootName() string {
	return "post"
}
