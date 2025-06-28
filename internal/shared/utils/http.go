package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIntQuery(c *gin.Context, key string, fallback int) int {
	valStr := c.Query(key)
	if valStr == "" {
		return fallback
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return fallback
	}
	return val
}
