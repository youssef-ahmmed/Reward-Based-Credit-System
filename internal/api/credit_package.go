package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCreditPackageRoutes(rg *gin.RouterGroup, handler *handler.CreditPackageHandler) {
	creditPackages := rg.Group("/credit-packages")

	creditPackages.GET("", handler.GetAllCreditPackages)
	creditPackages.GET("/:id", handler.GetCreditPackageByID)
	creditPackages.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.CreateCreditPackage)
	creditPackages.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.UpdateCreditPackages)
	creditPackages.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), handler.DeleteCreditPackage)
}
