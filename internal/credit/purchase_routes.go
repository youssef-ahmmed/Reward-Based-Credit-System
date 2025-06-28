package credit

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterPurchaseRoutes(rg *gin.RouterGroup, h *Handler) {
	purchases := rg.Group("/purchases", middleware.AuthMiddleware())

	purchases.POST("", h.CreatePurchase)
	purchases.GET("", h.GetUserPurchases)
	purchases.GET("/:id", h.GetPurchaseByID)
}
