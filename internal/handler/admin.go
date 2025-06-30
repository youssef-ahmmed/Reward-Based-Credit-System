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

// GetAdminDashboard godoc
// @Summary Get admin dashboard stats
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Success 200 {object} types.DashboardStatsResponse
// @Failure 500 {object} map[string]string
// @Router /admin/dashboard [get]
func (h *AdminHandler) GetAdminDashboard(c *gin.Context) {
	stats, err := h.service.GetAdminDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load dashboard stats"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stats": stats})
}

// GetAllUsers godoc
// @Summary Get all users with pagination, filtering and sorting
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param search query string false "Search keyword"
// @Param sort_by query string false "Sort by field"
// @Param sort_order query string false "Sort order (asc/desc)"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /admin/users [get]
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

// GetAllPurchases godoc
// @Summary Get all purchases with filters
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "Purchase status"
// @Param date_from query string false "Start date"
// @Param date_to query string false "End date"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /admin/purchases [get]
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

// GetAllRedemptions godoc
// @Summary Get all redemptions with filters
// @Tags Admin
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param status query string false "Redemption status"
// @Param date_from query string false "Start date"
// @Param date_to query string false "End date"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /admin/redemptions [get]
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

// UpdateRedemptionStatus godoc
// @Summary Update status of a redemption
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "Redemption ID"
// @Param request body types.UpdateRedemptionStatusRequest true "Status update payload"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /admin/redemptions/{id}/status [put]
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

// ManageUserCredits godoc
// @Summary Add or subtract user credits
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body types.ManageCreditsRequest true "Credits action payload"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /admin/users/{id}/credits [post]
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

// ManageUserPoints godoc
// @Summary Add or subtract user points
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body types.ManagePointsRequest true "Points action payload"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /admin/users/{id}/points [post]
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

// ModerateUser godoc
// @Summary Update user status (active, banned, suspended)
// @Tags Admin
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body types.ModerateUserRequest true "Moderate user request"
// @Success 200 {object} map[string]string
// @Failure 400,404,500 {object} map[string]string
// @Router /admin/users/{id}/status [put]
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
