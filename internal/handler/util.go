package handler

import (
	"strconv"
	"strings"
)

func parseInt(val string, def int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return def
	}
	return i
}

func parseBoolPtr(val string) *bool {
	if strings.ToLower(val) == "true" {
		b := true
		return &b
	} else if strings.ToLower(val) == "false" {
		b := false
		return &b
	}
	return nil
}
