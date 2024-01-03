package router

import (
	"github.com/gin-gonic/gin"
	"inlove-app-server/controllers"
	"inlove-app-server/controllers/posts"
)

// Router returns a gin router with all the routes defined.
func Router() *gin.Engine {
	router := gin.Default()

	definedControllers := []controllers.IResourceController{
		new(posts.PostController),
	}

	api := router.Group("/api")
	{
		for _, controller := range definedControllers {
			api.POST("/"+controller.RootName(), controller.Create)
			api.GET("/"+controller.RootName(), controller.List)
		}
	}
	return router
}
