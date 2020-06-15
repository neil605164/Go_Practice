package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// MyHandler01 測試 mock 使用
func (h *Handler) MyHandler01(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"a": 1,
	})
}
