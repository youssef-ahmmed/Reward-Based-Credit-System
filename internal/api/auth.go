package api

import (
	"Start/internal/handler"
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *handler.AuthHandler) {
	auth := rg.Group("/auth")

	auth.POST("/signup", handler.SignUp)
	auth.POST("/login", handler.Login)
	auth.POST("/refresh", handler.RefreshToken)
	auth.PUT("/change-password", middleware.AuthMiddleware(), handler.ChangePassword)
}
