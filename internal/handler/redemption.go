package handler

import (
	"Start/internal/service"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type RedemptionHandler struct {
	service service.RedemptionService
}

func NewRedemptionHandler(service service.RedemptionService) *RedemptionHandler {
	return &RedemptionHandler{service}
}

func (h *RedemptionHandler) CreateRedemption(c *gin.Context) {
	var req types.CreateRedemptionRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	userID := c.GetString("userId")

	resp, err := h.service.CreateRedemption(userID, req)
	if err != nil {
		switch err.Error() {
		case "product not found", "user wallet not found":
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "product is not available for redemption", "insufficient points", "invalid quantity":
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case "insufficient stock":
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Redemption failed"})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{"redemption": resp})
}

func (h *RedemptionHandler) GetUserRedemptions(c *gin.Context) {
	userID := c.GetString("userId")
	page := parseInt(c.Query("page"), 1)
	limit := parseInt(c.Query("limit"), 20)

	redemptions, total, err := h.service.GetUserRedemptions(userID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch redemptions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"redemptions": redemptions,
		"pagination": gin.H{
			"currentPage":  page,
			"itemsPerPage": limit,
			"totalItems":   total,
			"totalPages":   int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func (h *RedemptionHandler) GetRedemptionByID(c *gin.Context) {
	userID := c.GetString("userId")
	id := c.Param("id")

	r, err := h.service.GetRedemptionByID(userID, id)
	if err != nil {
		switch err.Error() {
		case "not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Redemption not found"})
		case "unauthorized":
			c.JSON(http.StatusForbidden, gin.H{"error": "Not allowed to access this redemption"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch redemption"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"redemption": r})
}
