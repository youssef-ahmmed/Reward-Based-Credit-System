package user

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterWalletRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/wallets", middleware.AuthMiddleware(), handler.GetWallet)
}
