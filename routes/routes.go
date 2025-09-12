package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafinhacuri/api-expo-go/controllers"
	"github.com/rafinhacuri/api-expo-go/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "Not Found",
		})
	})

	server.POST("/auth", controllers.Auth)

	api := server.Group("/api", middleware.Authenticate)

	api.POST("/logout", controllers.Logout)

	api.GET("/users", controllers.GetUsers)
	api.POST("/user", controllers.InsertUser)
	api.GET("/user", controllers.GetUser)
	api.DELETE("/user", controllers.DeleteUser)
	api.PUT("/user", controllers.UpdateUser)

	api.POST("/task", controllers.InsertTask)
	api.GET("/tasks", controllers.GetTasks)
	api.DELETE("/task", controllers.DeleteTask)
	api.PATCH("/task", controllers.CheckTask)
	api.PUT("/task", controllers.UpdateTask)
}
