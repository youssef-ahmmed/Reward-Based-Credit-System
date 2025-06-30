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

// RecommendProducts godoc
// @Summary Get AI-based product recommendations
// @Description Returns product recommendations based on user preferences, point balance, and context
// @Tags AI
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body types.RecommendationRequest true "Recommendation Input"
// @Success 200 {object} types.RecommendationResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /ai/recommendations [post]
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
