package handler

import (
	"Go_Practice/app/global/structs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// MyHandler01 測試呼叫 api 使用 mock 做測試
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

// MyHandler02 測試呼叫 redis 使用 mock 做測試
func (h *Handler) MyHandler02(c *gin.Context) {
	key := c.Query("key")

	time.Sleep(time.Second * 1)

	value, err := h.BInter.GetRedis(key)
	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"res": err.Error(),
		})
		return
	}

	num, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"res": err.Error(),
		})
		return
	}

	// 進行加總
	res := num + 10

	// 回傳結果
	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

// MyHandler03 測試呼叫 DB 使用 mock 做測試
func (h *Handler) MyHandler03(c *gin.Context) {
	var err error

	// 取參數
	req := structs.RawData{}
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})
		return
	}

	if err = h.BInter.StoreDBInfo(req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})
		return
	}
	// 回傳結果
	c.JSON(http.StatusOK, gin.H{
		"res": req,
	})
}
