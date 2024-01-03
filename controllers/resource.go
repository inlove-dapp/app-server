package controllers

import "github.com/gin-gonic/gin"

type IResourceController interface {
	// Create a new resource
	Create(c *gin.Context)
	// List all resources
	List(c *gin.Context)
	// Get resource by id
	Get(c *gin.Context)
	// Update resource by id
	Update(c *gin.Context)
	// RootName returns the name of the resource. Will be used to generate the routes.
	RootName() string
}
