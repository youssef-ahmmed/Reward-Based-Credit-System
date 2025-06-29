package app

import (
	"Start/internal/api"
	"Start/internal/handler"
	"Start/internal/repository"
	"Start/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserModule(rg *gin.RouterGroup, db *gorm.DB) {
	repo := repository.NewRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)
	api.RegisterUserRoutes(rg, h)
}
