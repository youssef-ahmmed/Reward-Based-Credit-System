package handler

import (
	"Start/internal/service"
	"Start/internal/shared/utils"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseHandler struct {
	service service.PurchaseService
}

func NewPurchaseHandler(service service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{service}
}

func (h *PurchaseHandler) CreatePurchase(c *gin.Context) {
	var req types.CreatePurchaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	userID := c.GetString("userId")
	resp, err := h.service.CreatePurchase(userID, req)
	if err != nil {
		if err.Error() == "package not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else if err.Error() == "payment failed" {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"purchase": resp})
}

func (h *PurchaseHandler) GetUserPurchases(c *gin.Context) {
	userID := c.GetString("userId")
	status := c.Query("status")
	page := utils.ParseIntQuery(c, "page", 1)
	limit := utils.ParseIntQuery(c, "limit", 20)

	purchases, meta, err := h.service.GetUserPurchases(userID, status, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch purchases"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"purchases": purchases, "pagination": meta})
}

func (h *PurchaseHandler) GetPurchaseByID(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetString("userId")

	resp, err := h.service.GetPurchaseByID(userID, id)
	if err != nil {
		if err.Error() == "not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Purchase not found"})
		} else if err.Error() == "forbidden" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"purchase": resp})
}
