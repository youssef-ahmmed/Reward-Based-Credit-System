package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPurchaseModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewPurchaseService(repo)
	h := handler.NewPurchaseHandler(svc)
	api.RegisterPurchaseRoutes(rg, h)
}
