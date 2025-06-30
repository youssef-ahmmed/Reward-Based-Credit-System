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

// CreatePurchase godoc
// @Summary Create a new purchase
// @Description User purchases a credit package. Requires authentication.
// @Tags Purchases
// @Accept json
// @Produce json
// @Param request body types.CreatePurchaseRequest true "Purchase input"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string "Invalid data"
// @Failure 402 {object} map[string]string "Payment failed"
// @Failure 404 {object} map[string]string "Package not found"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /purchases [post]
// @Security BearerAuth
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

// GetUserPurchases godoc
// @Summary Get all purchases by the authenticated user
// @Description Returns paginated purchases for a user with optional status filter
// @Tags Purchases
// @Produce json
// @Param status query string false "Filter by status (e.g. completed, pending)"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /purchases [get]
// @Security BearerAuth
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

// GetPurchaseByID godoc
// @Summary Get details of a specific purchase
// @Description Returns a purchase if it belongs to the authenticated user
// @Tags Purchases
// @Produce json
// @Param id path string true "Purchase ID"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]string "Access denied"
// @Failure 404 {object} map[string]string "Not found"
// @Failure 500 {object} map[string]string "Internal error"
// @Router /purchases/{id} [get]
// @Security BearerAuth
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
