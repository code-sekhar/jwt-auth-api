package middleware

import (
	"jwt-auth-api/config"
	"jwt-auth-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleware(permissionName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID := c.MustGet("role_id").(uint)
		var role models.Role
		err := config.DB.Preload("Permission").First(&role, roleID).Error
		if err != nil {
			c.JSON(
				http.StatusForbidden,
				gin.H{"message": "Role not found"},
			)
			c.Abort()
			return
		}
		for _, p := range role.Permission {
			if p.Name == permissionName {
				c.Next()
				return
			}
		}
		c.JSON(
			http.StatusForbidden,
			gin.H{"message": "Permission denied"},
		)
		c.Abort()
	}
}
