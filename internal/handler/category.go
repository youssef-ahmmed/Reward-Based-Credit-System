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

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteCategory(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
