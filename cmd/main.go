package main

import (
	"Start/internal/shared/database"
	"Start/internal/user"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Init DB
	db := database.InitDB()

	// Group versioned API routes
	api := r.Group("/api")

	// Register all user-related routes dynamically
	user.RegisterUserRoutes(api.Group("/users"), db)

	// Run server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("❌ App crashed: %v", err)
		return
	}
}
