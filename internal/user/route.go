package user

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	rg.POST("/signup", handler.SignUp)
	rg.POST("/login", handler.Login)
	rg.POST("/refresh", handler.RefreshToken)
	rg.PUT("/change-password", middleware.AuthMiddleware(), handler.ChangePassword)
}
