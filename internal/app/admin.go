package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAdminModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewAdminService(repo)
	h := handler.NewAdminHandler(svc)
	api.RegisterAdminRoutes(rg, h)
}
