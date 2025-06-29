package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterWalletRoutes(rg *gin.RouterGroup, handler *handler.WalletHandler) {
	rg.GET("/wallets", middleware.AuthMiddleware(), handler.GetWallet)
}
