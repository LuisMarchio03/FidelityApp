package router

import (
	docs "github.com/LuisMarchio03/acim-backend/docs"
	cardFidelityHandler "github.com/LuisMarchio03/acim-backend/handler/cardFidelity"
	companyHandler "github.com/LuisMarchio03/acim-backend/handler/company"
	pointHandler "github.com/LuisMarchio03/acim-backend/handler/point"
	userHandler "github.com/LuisMarchio03/acim-backend/handler/user"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(r *gin.Engine) {
	// Initialize the handlers
	companyHandler.InitializeHandler()
	userHandler.InitializeHandler()
	pointHandler.InitializeHandler()
	cardFidelityHandler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	// Initialize the routes
	v1 := r.Group("/api/v1")

	// Company
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
	}

	// User
	{
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

	// Point
	{
		v1.GET("/points", func(ctx *gin.Context) {
			pointHandler.ListPointsHandler(ctx)
		})
		v1.POST("/point", func(ctx *gin.Context) {
			pointHandler.CreatePointHandler(ctx)
		})
		v1.DELETE("/point", func(ctx *gin.Context) {
			pointHandler.DeletePointHandler(ctx)
		})
		v1.GET("/point", func(ctx *gin.Context) {
			pointHandler.ShowPointHandler(ctx)
		})
		v1.PUT("/point", func(ctx *gin.Context) {
			pointHandler.UpdatePointHandler(ctx)
		})
	}

	// CardFidelity
	{
		v1.GET("/cardFidelities", func(ctx *gin.Context) {
			cardFidelityHandler.ListCardFidelitiesHandler(ctx)
		})
		v1.POST("/cardFidelity", func(ctx *gin.Context) {
			cardFidelityHandler.CreateCardFidelityHandler(ctx)
		})
		v1.DELETE("/cardFidelity", func(ctx *gin.Context) {
			cardFidelityHandler.DeleteCardFidelityHandler(ctx)
		})
		v1.GET("/cardFidelity", func(ctx *gin.Context) {
			cardFidelityHandler.ShowCardFidelityHandler(ctx)
		})
		v1.PUT("/cardFidelity", func(ctx *gin.Context) {
			cardFidelityHandler.UpdateCardFidelityHandler(ctx)
		})
	}

	// Intialize Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
	))
}
