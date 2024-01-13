package router

import (
	"github.com/gin-gonic/gin"
	"inlove-app-server/clients"
	"inlove-app-server/controllers"
	"inlove-app-server/controllers/itinerary_report"
	"inlove-app-server/controllers/posts"
	"inlove-app-server/controllers/user"
	"inlove-app-server/db"
	"inlove-app-server/middlewares"
)

// Router returns a gin router with all the routes defined.
func Router() *gin.Engine {
	router := gin.Default()
	client := db.GetDB()
	prismaErrorClient := clients.NewPrismaErrorClient()

	definedControllers := []controllers.IResourceController{
		posts.NewPostController(client),
		itinerary_report.NewItineraryReportController(client, prismaErrorClient),
		user.NewUserController(client, prismaErrorClient),
	}

	api := router.Group("/api").Use(middlewares.JWTAuthMiddleware())
	{
		for _, controller := range definedControllers {
			api.POST("/"+controller.RootName(), controller.Create)
			api.GET("/"+controller.RootName(), controller.List)
		}
	}
	return router
}
