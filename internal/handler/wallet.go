package handler

import (
	"Start/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WalletHandler struct {
	service service.WalletService
}

func NewWalletHandler(service service.WalletService) *WalletHandler {
	return &WalletHandler{service}
}

func (h *WalletHandler) GetWallet(c *gin.Context) {
	userID := c.GetString("userId")

	wallet, err := h.service.GetWallet(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch wallet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wallet": gin.H{
			"user_id":         wallet.UserID,
			"points_balance":  wallet.PointsBalance,
			"credits_balance": wallet.CreditsBalance,
			"updated_at":      wallet.UpdatedAt.Format(time.RFC3339),
		},
	})
}
