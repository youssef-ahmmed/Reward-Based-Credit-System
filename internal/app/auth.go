package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAuthModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewAuthService(repo)
	h := handler.NewAuthHandler(svc)
	api.RegisterAuthRoutes(rg, h)
}
