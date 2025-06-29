package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCreditPackageModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewCreditPackageService(repo)
	h := handler.NewCreditPackageHandler(svc)
	api.RegisterCreditPackageRoutes(rg, h)
}
