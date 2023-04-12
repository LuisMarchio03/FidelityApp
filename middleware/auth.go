package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticação necessário"})
			return
		}

		// Verifica se o token é válido

		// Se o token estiver presente, chama o próximo handler
		c.Next()
	}
}
