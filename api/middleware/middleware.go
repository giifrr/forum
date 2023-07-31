package middleware

import (
	"net/http"

	"github.com/giifrr/forum/api/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := auth.TokenValid(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"status": http.StatusUnauthorized,
				"message": "Unauthorized",
			})

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

// This enables us interact with the React Frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
