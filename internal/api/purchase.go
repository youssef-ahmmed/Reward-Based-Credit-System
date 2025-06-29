package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(rg *gin.RouterGroup, handler *handler.PurchaseHandler) {
	purchases := rg.Group("/purchases", middleware.AuthMiddleware())

	purchases.POST("", handler.CreatePurchase)
	purchases.GET("", handler.GetUserPurchases)
	purchases.GET("/:id", handler.GetPurchaseByID)
}
