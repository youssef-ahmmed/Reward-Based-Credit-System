package main

import (
	_ "Start/docs"
	"Start/internal/app"
	"Start/internal/migration"
	"Start/internal/shared/database"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Reward System API
// @version 1.0
// @description API documentation for Reward-Based Credit System

// @host localhost:8080
// @BasePath /api
func main() {
	r := gin.Default()

	db := database.GetDB()

	if err := migration.AutoMigrate(db); err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.RegisterModules(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("App crashed: %v", err)
		return
	}
}
