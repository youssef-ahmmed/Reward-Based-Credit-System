package handler

import (
	"Start/internal/service"
	"Start/internal/types"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AIHandler struct {
	service service.AIService
}

func NewAIHandler(service service.AIService) *AIHandler {
	return &AIHandler{service}
}

func (h *AIHandler) RecommendProducts(c *gin.Context) {
	var req types.RecommendationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := h.service.RecommendProducts(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate recommendations"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
