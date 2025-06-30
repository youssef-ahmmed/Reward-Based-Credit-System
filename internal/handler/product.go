package handler

import (
	"Start/internal/service"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service}
}

// GetAllProducts godoc
// @Summary Get all products
// @Description Retrieve all products with optional filters and pagination
// @Tags Products
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Param sort_by query string false "Sort by field (e.g., created_at, reward_points)"
// @Param sort_order query string false "Sort order (asc or desc)"
// @Param category_id query string false "Filter by category ID"
// @Param is_offer query bool false "Filter by offer status"
// @Param min_points query int false "Minimum reward points"
// @Param max_points query int false "Maximum reward points"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /products [get]
func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	page := parseInt(c.Query("page"), 1)
	limit := parseInt(c.Query("limit"), 20)
	sortBy := c.Query("sort_by")
	sortOrder := c.Query("sort_order")

	filters := types.ProductFilters{
		CategoryID: c.Query("category_id"),
		IsOffer:    parseBoolPtr(c.Query("is_offer")),
		MinPoints:  parseInt(c.Query("min_points"), 0),
		MaxPoints:  parseInt(c.Query("max_points"), 0),
	}

	products, meta, err := h.service.GetAllProducts(filters, page, limit, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products, "pagination": meta})
}

// SearchProducts godoc
// @Summary Search products
// @Description Search products by query and optional filters
// @Tags Products
// @Produce json
// @Param query query string true "Search query"
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(20)
// @Param category_id query string false "Filter by category ID"
// @Param is_offer query bool false "Filter by offer status"
// @Param min_points query int false "Minimum reward points"
// @Param max_points query int false "Maximum reward points"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/search [get]
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}
	page := parseInt(c.Query("page"), 1)
	limit := parseInt(c.Query("limit"), 20)

	filters := types.ProductFilters{
		CategoryID: c.Query("category_id"),
		IsOffer:    parseBoolPtr(c.Query("is_offer")),
		MinPoints:  parseInt(c.Query("min_points"), 0),
		MaxPoints:  parseInt(c.Query("max_points"), 0),
	}

	products, meta, err := h.service.SearchProducts(query, filters, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Search failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products, "pagination": meta})
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Admin can create a new redeemable product
// @Tags Products
// @Accept json
// @Produce json
// @Param request body types.CreateProductRequest true "Product data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products [post]
// @Security BearerAuth
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req types.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	product, err := h.service.CreateProduct(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creation failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Admin can update a product by ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param request body types.UpdateProductRequest true "Updated product data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /products/{id} [put]
// @Security BearerAuth
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req types.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	product, err := h.service.UpdateProduct(id, &req)
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UpdateCreditPackage failed"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Admin can delete a product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /products/{id} [delete]
// @Security BearerAuth
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
