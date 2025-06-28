package credit

import (
	"Start/internal/shared/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllCreditPackages(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))

	var activeFilter *bool
	if active := c.Query("active"); active != "" {
		val := active == "true"
		activeFilter = &val
	}

	pkgs, pagination, err := h.service.GetAllCreditPackages(page, limit, activeFilter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch packages"})
		return
	}

	c.JSON(http.StatusOK, PaginatedResponse{
		Packages:   pkgs,
		Pagination: pagination,
	})
}

func (h *Handler) GetCreditPackageByID(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.service.GetCreditPackageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"package": pkg})
}

func (h *Handler) CreateCreditPackage(c *gin.Context) {
	var req CreateCreditPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	pkg, err := h.service.CreateCreditPackage(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create package"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "CreditPackage created successfully",
		"package": pkg,
	})
}

func (h *Handler) UpdateCreditPackages(c *gin.Context) {
	id := c.Param("id")
	var req UpdateCreditPackageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	pkg, err := h.service.UpdateCreditPackages(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "CreditPackage updated successfully",
		"package": pkg,
	})
}

func (h *Handler) DeleteCreditPackage(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCreditPackage(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) CreatePurchase(c *gin.Context) {
	var req CreatePurchaseRequest
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

func (h *Handler) GetUserPurchases(c *gin.Context) {
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

func (h *Handler) GetPurchaseByID(c *gin.Context) {
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
