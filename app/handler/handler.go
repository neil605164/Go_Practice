package handler

import (
	"Go_Practice/app/global/structs"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// MyHandler01 測試 mock 使用
func (h *Handler) MyHandler01(c *gin.Context) {

	// 取參數
	req := structs.UrlQuery{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})

		return
	}

	time.Sleep(1 * time.Second)

	// 加法運算
	res := h.BInter.Api(req.Num1, req.Num2)
	c.JSON(http.StatusOK, gin.H{
		"a": res,
	})
}
