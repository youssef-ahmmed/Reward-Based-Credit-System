package product

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRedemptionRoutes(rg *gin.RouterGroup, h *Handler) {
	redemption := rg.Group("/redemptions")

	redemption.POST("", middleware.AuthMiddleware(), h.CreateRedemption)
	redemption.GET("", middleware.AuthMiddleware(), h.GetUserRedemptions)
	redemption.GET("/:id", middleware.AuthMiddleware(), h.GetRedemptionByID)
}
