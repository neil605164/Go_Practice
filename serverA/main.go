package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/service", Service01)

	r.Run(":7897")
}

type formData struct {
	A int16 `form:"a" binding:"required"`
	B int16 `form:"b"`
}

func Service01(c *gin.Context) {
	// 取參數
	req := formData{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"res": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"res": req.A + req.B,
	})
}
