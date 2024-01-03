package router

import (
	"github.com/gin-gonic/gin"
	"inlove-app-server/controllers"
)

func Router() *gin.Engine {
	router := gin.Default()

	authController := new(controllers.AuthenticationController)

	api := router.Group("/api")
	{
		api.POST("/register", authController.Register)
	}

	return router
}
