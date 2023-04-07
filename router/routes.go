package router

import (
	"github.com/LuisMarchio03/acim-backend/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// Initialize the handlers
	handler.InitializeHandler()

	// Initialize the routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/companies", func(ctx *gin.Context) {
			handler.ListCompaniesHandler(ctx)
		})
	}
}
