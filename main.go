package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", MyHandler)
	r.POST("/pong", MyHandler02)

	r.Run(":9698")
}

func MyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"a": 1,
	})
}

func MyHandler02(c *gin.Context) {

	val := c.PostForm("a")

	c.JSON(http.StatusOK, gin.H{
		"a": val,
	})
}
