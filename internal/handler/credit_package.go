package handler

import (
	"Start/internal/service"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CreditPackageHandler struct {
	service service.CreditPackageService
}

func NewCreditPackageHandler(service service.CreditPackageService) *CreditPackageHandler {
	return &CreditPackageHandler{service}
}

func (h *CreditPackageHandler) GetAllCreditPackages(c *gin.Context) {
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

	c.JSON(http.StatusOK, types.PaginatedResponse{
		Packages:   pkgs,
		Pagination: pagination,
	})
}

func (h *CreditPackageHandler) GetCreditPackageByID(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.service.GetCreditPackageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"package": pkg})
}

func (h *CreditPackageHandler) CreateCreditPackage(c *gin.Context) {
	var req types.CreateCreditPackageRequest
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

func (h *CreditPackageHandler) UpdateCreditPackages(c *gin.Context) {
	id := c.Param("id")
	var req types.UpdateCreditPackageRequest
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

func (h *CreditPackageHandler) DeleteCreditPackage(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCreditPackage(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
