package user

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterProfileRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.GET("/profile", middleware.AuthMiddleware(), handler.GetProfile)
	rg.PUT("/profile", middleware.AuthMiddleware(), handler.UpdateProfile)
}
