package main

import (
	"taskflow/internal/db"
	"taskflow/internal/handlers"
          "taskflow/internal/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to database
	db.Connect()

	// Create Gin router
	r := gin.Default()

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "TaskFlow API running"})
	})

	// Public routes
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	// 🟢 ADD THIS BLOCK
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	// Test protected route
	protected.GET("/protected", func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		c.JSON(200, gin.H{"message": "Access granted", "user_id": userID})
	})

	
	protected.POST("/projects", handlers.CreateProject)
	protected.GET("/projects", handlers.GetProjects)

	protected.POST("/projects/:id/tasks", handlers.CreateTask)
protected.GET("/projects/:id/tasks", handlers.GetTasks)
protected.PUT("/tasks/:id", handlers.UpdateTask)
protected.DELETE("/tasks/:id", handlers.DeleteTask)

	// Start server
	r.Run(":5000")
}