package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, handler *handler.UserHandler) {
	rg.GET("/profile", middleware.AuthMiddleware(), handler.GetProfile)
	rg.PUT("/profile", middleware.AuthMiddleware(), handler.UpdateProfile)
}
