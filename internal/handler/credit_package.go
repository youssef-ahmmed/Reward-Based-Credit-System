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

// GetAllCreditPackages godoc
// @Summary Get all credit packages
// @Description Fetch all available credit packages with pagination and optional active filter
// @Tags Credit Packages
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Param active query bool false "Filter by active status"
// @Success 200 {object} types.PaginatedResponse
// @Failure 500 {object} map[string]string
// @Router /credit-packages [get]
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

// GetCreditPackageByID godoc
// @Summary Get credit package by ID
// @Description Retrieve a single credit package by its ID
// @Tags Credit Packages
// @Produce json
// @Param id path string true "CreditPackage ID"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Router /credit-packages/{id} [get]
func (h *CreditPackageHandler) GetCreditPackageByID(c *gin.Context) {
	id := c.Param("id")
	pkg, err := h.service.GetCreditPackageByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"package": pkg})
}

// CreateCreditPackage godoc
// @Summary Create a new credit package
// @Description Admin can create a new credit package with points and credits
// @Tags Credit Packages
// @Accept json
// @Produce json
// @Param request body types.CreateCreditPackageRequest true "Credit Package creation data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /credit-packages [post]
// @Security BearerAuth
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

// UpdateCreditPackages godoc
// @Summary Update an existing credit package
// @Description Admin can update a credit package by ID
// @Tags Credit Packages
// @Accept json
// @Produce json
// @Param id path string true "CreditPackage ID"
// @Param request body types.UpdateCreditPackageRequest true "Credit Package update data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /credit-packages/{id} [put]
// @Security BearerAuth
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

// DeleteCreditPackage godoc
// @Summary Delete a credit package
// @Description Admin can delete a credit package by ID
// @Tags Credit Packages
// @Produce json
// @Param id path string true "CreditPackage ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /credit-packages/{id} [delete]
// @Security BearerAuth
func (h *CreditPackageHandler) DeleteCreditPackage(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCreditPackage(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CreditPackage not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
