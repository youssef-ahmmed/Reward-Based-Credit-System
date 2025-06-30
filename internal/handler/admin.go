package handler

import (
	"Start/internal/service"
	"Start/internal/shared/utils"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

type AdminHandler struct {
	service service.AdminService
}

func NewAdminHandler(service service.AdminService) *AdminHandler {
	return &AdminHandler{service}
}

func (h *AdminHandler) GetAdminDashboard(c *gin.Context) {
	stats, err := h.service.GetAdminDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load dashboard stats"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	page := utils.ParseIntQuery(c, c.Query("page"), 1)
	limit := utils.ParseIntQuery(c, c.Query("limit"), 20)
	search := c.Query("search")
	sortBy := c.DefaultQuery("sort_by", "created_at")
	sortOrder := c.DefaultQuery("sort_order", "desc")

	users, total, err := h.service.GetAllUsers(page, limit, search, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"pagination": gin.H{
			"currentPage":  page,
			"itemsPerPage": limit,
			"totalItems":   total,
			"totalPages":   int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func (h *AdminHandler) GetAllPurchases(c *gin.Context) {
	page := utils.ParseIntQuery(c, c.Query("page"), 1)
	limit := utils.ParseIntQuery(c, c.Query("limit"), 20)
	status := c.Query("status")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	purchases, total, err := h.service.GetAllPurchases(page, limit, status, dateFrom, dateTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch purchases"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"purchases": purchases,
		"pagination": gin.H{
			"currentPage":  page,
			"itemsPerPage": limit,
			"totalItems":   total,
			"totalPages":   int(math.Ceil(float64(total) / float64(limit))),
		},
	})
}

func (h *AdminHandler) GetAllRedemptions(c *gin.Context) {
	page := utils.ParseIntQuery(c, c.Query("page"), 1)
	limit := utils.ParseIntQuery(c, c.Query("limit"), 20)
	status := c.Query("status")
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	redemptions, total, err := h.service.GetAllRedemptions(page, limit, status, dateFrom, dateTo)
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

func (h *AdminHandler) UpdateRedemptionStatus(c *gin.Context) {
	var req types.UpdateRedemptionStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status payload"})
		return
	}
	id := c.Param("id")

	if err := h.service.UpdateRedemptionStatus(id, req.Status); err != nil {
		switch err.Error() {
		case "not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "Redemption not found"})
		case "invalid status":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update status"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Status updated successfully"})
}

func (h *AdminHandler) ManageUserCredits(c *gin.Context) {
	userID := c.Param("id")

	var req types.ManageCreditsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.service.ManageUserCredits(userID, req.Action, req.Amount)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "invalid action":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credit action"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update credits"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credits updated successfully"})
}

func (h *AdminHandler) ManageUserPoints(c *gin.Context) {
	userID := c.Param("id")

	var req types.ManagePointsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.service.ManageUserPoints(userID, req.Action, req.Amount)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "invalid action":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid points action"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update points"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Points updated successfully"})
}

func (h *AdminHandler) ModerateUser(c *gin.Context) {
	userID := c.Param("id")

	var req types.ModerateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := h.service.UpdateUserStatus(userID, req.Status)
	if err != nil {
		switch err.Error() {
		case "user not found":
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		case "invalid status":
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User status updated"})
}
