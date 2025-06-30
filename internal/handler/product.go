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

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
