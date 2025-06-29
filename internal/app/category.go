package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCategoryModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewCategoryService(repo)
	h := handler.NewCategoryHandler(svc)
	api.RegisterCategoryRoutes(rg, h)
}
