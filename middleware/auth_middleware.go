package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"message": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(
				os.Getenv("JWT_SECRET"),
			), nil
		},
		)
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"message": "Invalid token"})
			c.Abort()
			return
		}
		claims := token.Claims.(jwt.MapClaims)

		c.Set("user_id", claims["user_id"])
		c.Set(
			"role_id",
			uint(claims["role_id"].(float64)),
		)
		//c.Set("email", claims["email"])
		c.Next()
	}
}
