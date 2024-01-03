package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"inlove-app-server/db"
	"net/http"
)

var (
	prismaClient = db.GetDB()
	contextB     = context.Background()
)

type IAuthenticationController interface {
	Register(c *gin.Context)
}

type AuthenticationController struct {
}

func (ac AuthenticationController) Register(c *gin.Context) {
	posts, err := prismaClient.Post.FindMany().Exec(contextB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, posts)
}
