package credit

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCreditRoutes(rg *gin.RouterGroup, h *Handler) {
	creditPackages := rg.Group("/credit-packages")

	creditPackages.GET("", h.GetAllCreditPackages)
	creditPackages.GET("/:id", h.GetCreditPackageByID)
	creditPackages.POST("", middleware.AuthMiddleware(), middleware.AdminMiddleware(), h.CreateCreditPackage)
	creditPackages.PUT("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), h.UpdateCreditPackages)
	creditPackages.DELETE("/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), h.DeleteCreditPackage)
}
