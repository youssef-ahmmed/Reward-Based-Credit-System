package credit

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCreditRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/credit-packages", handler.GetAllCreditPackages)
	rg.GET("/credit-packages/:id", handler.GetCreditPackageByID)

	adminRoutes := rg.Group("/credit-packages", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminRoutes.POST("", handler.CreateCreditPackage)
		adminRoutes.PUT("/:id", handler.UpdateCreditPackages)
		adminRoutes.DELETE("/:id", handler.DeleteCreditPackage)
	}
}
