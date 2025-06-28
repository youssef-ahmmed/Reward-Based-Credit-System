package product

import (
	"Start/internal/shared/database"
	"github.com/gin-gonic/gin"
)

func RegisterModule(rg *gin.RouterGroup) {
	db := database.GetDB()

	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	RegisterCategoryRoutes(rg, handler)
	RegisterProductRoutes(rg, handler)
}
