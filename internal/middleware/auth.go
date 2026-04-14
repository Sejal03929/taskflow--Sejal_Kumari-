package middleware

import (
	"net/http"
	"os"
	"strings"

	"taskflow/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid format"})
			c.Abort()
			return
		}

		// ✅ USE SAME CLAIMS STRUCT
		claims := &handlers.Claims{}

		token, err := jwt.ParseWithClaims(parts[1], claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// ✅ SET USER ID
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}