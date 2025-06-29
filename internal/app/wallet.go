package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterWalletModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewWalletService(repo)
	h := handler.NewWalletHandler(svc)
	api.RegisterWalletRoutes(rg, h)
}
