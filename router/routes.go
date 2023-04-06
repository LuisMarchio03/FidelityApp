package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	// Initialize the handlers
	// handler.InitializeHandler()

	// Initialize the routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
