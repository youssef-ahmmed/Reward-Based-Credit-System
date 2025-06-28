package product

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *Handler) {
	categories := rg.Group("/categories")

	categories.GET("", handler.GetAllCategories)
	categories.GET("/:id/details", handler.GetCategoryDetails)
	categories.POST("", middleware.AdminMiddleware(), handler.CreateCategory)
	categories.PUT("/:id", middleware.AdminMiddleware(), handler.UpdateCategory)
	categories.DELETE("/:id", middleware.AdminMiddleware(), handler.DeleteCategory)
}
