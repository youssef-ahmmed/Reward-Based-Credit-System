package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRedemptionModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewRedemptionService(repo)
	h := handler.NewRedemptionHandler(svc)
	api.RegisterRedemptionRoutes(rg, h)
}
