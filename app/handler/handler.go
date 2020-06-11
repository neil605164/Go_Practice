package handler

import (
	"Go_Practice/app/business"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyHandler01(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"a": 1,
	})
}

func MyHandler02(c *gin.Context) {

	val := c.PostForm("a")

	if val == "123" {
		val = "456"
	}

	time.Sleep(time.Second * 1)

	c.JSON(http.StatusOK, gin.H{
		"a": val,
	})
}

type rawData struct {
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

func MyHandler03(c *gin.Context) {
	req := rawData{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "Linda" {
		req.Sex = "male"
	}

	c.JSON(http.StatusOK, req)
}

type formData struct {
	A int8 `form:"a" binding:"required"`
	B int8 `form:"b"`
}

func MyHandler04(c *gin.Context) {
	// 取參數
	req := formData{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})

		return
	}

	// New interface 接口
	bus := business.NewBusiness()
	res := bus.Add(req.A, req.B)

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}

func MyHandler05(c *gin.Context) {
	// 取參數
	req := formData{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})

		return
	}

	// New interface 接口
	bus := business.NewBusiness()
	res := bus.Api(req.A, req.B)

	c.JSON(http.StatusOK, gin.H{
		"res": res,
	})
}
