package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterUserRoutes wires repository → service → handler → router
func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	rg.POST("/signup", handler.SignUp)
	// You can later add: rg.POST("/login", handler.Login) etc.
}
