package main

import (
	"jwt-auth-api/config"
	"jwt-auth-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
