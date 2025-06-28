package user

import (
	"Start/internal/shared/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *Handler) {
	rg.POST("/signup", handler.SignUp)
	rg.POST("/login", handler.Login)
	rg.POST("/refresh", handler.RefreshToken)
	rg.PUT("/change-password", middleware.AuthMiddleware(), handler.ChangePassword)
}
