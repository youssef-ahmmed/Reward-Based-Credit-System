package main

import (
	"Start/internal/app"
	"Start/internal/shared/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	db := database.GetDB()

	app.RegisterModules(r, db)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("App crashed: %v", err)
		return
	}
}
