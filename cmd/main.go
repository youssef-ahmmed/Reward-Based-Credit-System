package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	RegisterRoutes(r)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("❌ App crashed: %v", err)
		return
	}
}
