package main

import (
	"Start/internal/credit"
	"Start/internal/product"
	"Start/internal/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	user.RegisterModule(api.Group("/users"))
	credit.RegisterModule(api)
	product.RegisterModule(api)
}
