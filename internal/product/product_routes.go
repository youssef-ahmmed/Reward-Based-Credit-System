package product

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(rg *gin.RouterGroup, handler *Handler) {
	products := rg.Group("/products")

	products.GET("", handler.GetAllProducts)
	products.GET("/search", handler.SearchProducts)
	products.POST("", middleware.AdminMiddleware(), handler.CreateProduct)
	products.PUT("/:id", middleware.AdminMiddleware(), handler.UpdateProduct)
	products.DELETE("/:id", middleware.AdminMiddleware(), handler.DeleteProduct)
}
