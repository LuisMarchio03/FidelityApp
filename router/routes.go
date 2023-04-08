package router

import (
	docs "github.com/LuisMarchio03/acim-backend/docs"
	companyHandler "github.com/LuisMarchio03/acim-backend/handler/company"
	userHandler "github.com/LuisMarchio03/acim-backend/handler/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	// Initialize the handlers
	companyHandler.InitializeHandler()
	userHandler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Initialize the routes
	v1 := r.Group("/api/v1")
	{
		v1.POST("/company", func(ctx *gin.Context) {
			companyHandler.CreateCompanyHandler(ctx)
		})
		v1.GET("/companies", func(ctx *gin.Context) {
			companyHandler.ListCompaniesHandler(ctx)
		})
		v1.DELETE("/company", func(ctx *gin.Context) {
			companyHandler.DeleteCompanyHandler(ctx)
		})
		v1.GET("/company", func(ctx *gin.Context) {
			companyHandler.ShowCompanyHandler(ctx)
		})
		v1.PUT("/company", func(ctx *gin.Context) {
			companyHandler.UpdateCompanyHandler(ctx)
		})
		v1.POST("/user", func(ctx *gin.Context) {
			userHandler.CreateUserHandler(ctx)
		})
		v1.GET("/users", func(ctx *gin.Context) {
			userHandler.ListUsersHandler(ctx)
		})
		v1.DELETE("/user", func(ctx *gin.Context) {
			userHandler.DeleteUserHandler(ctx)
		})
		v1.GET("/user", func(ctx *gin.Context) {
			userHandler.ShowUserHandler(ctx)
		})
		v1.PUT("/user", func(ctx *gin.Context) {
			userHandler.UpdateUserHandler(ctx)
		})
	}

	// Intialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))
}
