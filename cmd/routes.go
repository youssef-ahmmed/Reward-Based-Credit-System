package main

import (
	"Start/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	user.RegisterModule(api.Group("/users"))
}
