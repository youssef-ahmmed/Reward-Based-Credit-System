package product

import (
	"Start/internal/bridge"
	"Start/internal/shared/database"
	"Start/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterModule(rg *gin.RouterGroup) {
	db := database.GetDB()

	repo := NewRepository(db)

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)

	walletBridge := bridge.NewWalletServiceBridge(userService)

	service := NewService(repo, walletBridge)
	handler := NewHandler(service)

	RegisterCategoryRoutes(rg, handler)
	RegisterProductRoutes(rg, handler)
	RegisterRedemptionRoutes(rg, handler)
}
