package user

import (
	"github.com/gin-gonic/gin"
	"inlove-app-server/clients"
	"inlove-app-server/constants/keys"
	"inlove-app-server/controllers"
	prisma "inlove-app-server/prisma/db"
	"inlove-app-server/types"
	"net/http"
)

type IUserController interface {
	controllers.IResourceController
}

type UserController struct {
	controllers.IResourceController
	prisma            *prisma.PrismaClient
	prismaErrorClient clients.IPrismaErrorClient
}

// NewUserController returns a new instance of UserController.
func NewUserController(prisma *prisma.PrismaClient, prismaErrorClient clients.IPrismaErrorClient) IUserController {
	return &UserController{
		prisma:            prisma,
		prismaErrorClient: prismaErrorClient,
	}
}

// Create is responsible for creating a new user in the database.
// It will be called when the user logs in for the first time.
func (uc *UserController) Create(c *gin.Context) {
	userData, _ := c.Get(keys.UserKey)
	user := userData.(types.User)
	createdUser, err := uc.prisma.User.CreateOne(
		prisma.User.Name.Set(user.Email),
		prisma.User.SupabaseUserID.Set(user.Id),
	).Exec(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": uc.prismaErrorClient.HandleError(err)})
		return
	}
	c.JSON(http.StatusOK, createdUser)
}

func (_ *UserController) RootName() string {
	return "user"
}
