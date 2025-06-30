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

// CreateRedemption godoc
// @Summary Redeem a product using points
// @Description Allows a user to redeem a product if they have enough points and stock is available
// @Tags Redemptions
// @Accept json
// @Produce json
// @Param request body types.CreateRedemptionRequest true "Redemption input"
// @Success 201 {object} map[string]interface{} "Redemption successful"
// @Failure 400 {object} map[string]string "Invalid request or insufficient points"
// @Failure 404 {object} map[string]string "Product or wallet not found"
// @Failure 409 {object} map[string]string "Insufficient stock"
// @Failure 500 {object} map[string]string "Redemption failed"
// @Router /redemptions [post]
// @Security BearerAuth
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

// GetUserRedemptions godoc
// @Summary Get all redemptions of the authenticated user
// @Description Returns paginated list of redemptions made by the user
// @Tags Redemptions
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{} "Paginated redemptions"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /redemptions [get]
// @Security BearerAuth
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

// GetRedemptionByID godoc
// @Summary Get details of a specific redemption
// @Description Returns redemption details if it belongs to the authenticated user
// @Tags Redemptions
// @Produce json
// @Param id path string true "Redemption ID"
// @Success 200 {object} map[string]interface{} "Redemption data"
// @Failure 403 {object} map[string]string "Not allowed"
// @Failure 404 {object} map[string]string "Redemption not found"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /redemptions/{id} [get]
// @Security BearerAuth
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
