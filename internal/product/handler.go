package product

import (
	"Start/internal/shared/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	page := parseInt(c.Query("page"), 1)
	limit := parseInt(c.Query("limit"), 20)
	sortBy := c.Query("sort_by")
	sortOrder := c.Query("sort_order")

	filters := ProductFilters{
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

func (h *Handler) SearchProducts(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query is required"})
		return
	}
	page := parseInt(c.Query("page"), 1)
	limit := parseInt(c.Query("limit"), 20)

	filters := ProductFilters{
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

func (h *Handler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest
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

func (h *Handler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	product, err := h.service.UpdateProduct(id, &req)
	if err != nil {
		if err.Error() == "product not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"product": product,
	})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.service.CreateCategory(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Creation failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":  "Category created successfully",
		"category": resp,
	})
}

func (h *Handler) GetAllCategories(c *gin.Context) {
	var parentID *string
	if p := c.Query("parent_id"); p != "" {
		parentID = &p
	}
	categories, err := h.service.GetAllCategories(parentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Fetch failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func (h *Handler) GetCategoryDetails(c *gin.Context) {
	id := c.Param("id")
	page := utils.ParseIntQuery(c, "page", 1)
	limit := utils.ParseIntQuery(c, "limit", 20)

	data, err := h.service.GetCategoryDetails(id, page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	category, err := h.service.UpdateCategory(id, &req)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Update failed"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Category updated successfully",
		"category": category,
	})
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
