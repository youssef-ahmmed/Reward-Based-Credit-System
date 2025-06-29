package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRedemptionRoutes(rg *gin.RouterGroup, handler *handler.RedemptionHandler) {
	redemption := rg.Group("/redemptions")

	redemption.POST("", middleware.AuthMiddleware(), handler.CreateRedemption)
	redemption.GET("", middleware.AuthMiddleware(), handler.GetUserRedemptions)
	redemption.GET("/:id", middleware.AuthMiddleware(), handler.GetRedemptionByID)
}
