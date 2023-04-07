package router

import (
	docs "github.com/LuisMarchio03/acim-backend/docs"
	"github.com/LuisMarchio03/acim-backend/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	// Initialize the handlers
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Initialize the routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/companies", func(ctx *gin.Context) {
			handler.ListCompaniesHandler(ctx)
		})
	}

	// Intialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))
}
