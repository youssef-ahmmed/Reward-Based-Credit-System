package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAIRoutes(rg *gin.RouterGroup, handler *handler.AIHandler) {
	ai := rg.Group("/ai", middleware.AuthMiddleware())
	ai.POST("/recommendations", handler.RecommendProducts)
}
