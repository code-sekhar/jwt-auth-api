package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := c.MustGet("role_id").(uint)

		for _, role := range allowedRoles {
			if roleID == role {
				c.Next()
				return
			}
		}
		c.JSON(
			http.StatusForbidden,
			gin.H{
				"message": "Access Denied",
			},
		)
		c.Abort()
	}
}
