package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	path := server.Group("/api")

	path.POST("/user", controllers.InsertUser)
}
