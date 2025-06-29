package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(rg *gin.RouterGroup, handler *handler.AdminHandler) {
	admin := rg.Group("/admin", middleware.AuthMiddleware(), middleware.AdminMiddleware())

	admin.GET("/dashboard", handler.GetAdminDashboard)
	admin.GET("/users", handler.GetAllUsers)
	admin.GET("/purchases", handler.GetAllPurchases)
	admin.GET("/redemptions", handler.GetAllRedemptions)

	admin.PUT("/redemptions/:id/status", handler.UpdateRedemptionStatus)
	admin.POST("/users/:id/credits", handler.ManageUserCredits)
	admin.POST("/users/:id/points", handler.ManageUserPoints) // todo
	admin.PUT("/users/:id/status", handler.ModerateUser)
}
