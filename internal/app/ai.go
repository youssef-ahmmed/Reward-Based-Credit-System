package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterAIModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewAIService(repo)
	h := handler.NewAIHandler(svc)
	api.RegisterAIRoutes(rg, h)
}
