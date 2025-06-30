package handler

import (
	"Start/internal/service"
	"Start/internal/shared/utils"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(service service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

// CreateCategory godoc
// @Summary Create new category
// @Description Create a new category with optional parent
// @Tags Categories
// @Accept json
// @Produce json
// @Param request body types.CreateCategoryRequest true "Category data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
// @Security BearerAuth
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var req types.CreateCategoryRequest
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

// GetAllCategories godoc
// @Summary List all categories
// @Description Fetch list of categories, optionally filtered by parent ID
// @Tags Categories
// @Produce json
// @Param parent_id query string false "Parent category ID"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
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

// GetCategoryDetails godoc
// @Summary Get category details and its products
// @Description Retrieve category details along with its paginated products
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} interface{}
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func (h *CategoryHandler) GetCategoryDetails(c *gin.Context) {
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

// UpdateCategory godoc
// @Summary Update category details
// @Description Update name or description of a category
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Param request body types.UpdateCategoryRequest true "Update data"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
// @Security BearerAuth
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req types.UpdateCategoryRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	category, err := h.service.UpdateCategory(id, &req)
	if err != nil {
		if err.Error() == "category not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "UpdateCreditPackage failed"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Category updated successfully",
		"category": category,
	})
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by ID
// @Tags Categories
// @Produce json
// @Param id path string true "Category ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [delete]
// @Security BearerAuth
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
