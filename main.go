package main

import (
	"Go_Practice/app/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hand01", handler.MyHandler01)
	r.POST("/hand02", handler.MyHandler02)
	r.POST("/hand03", handler.MyHandler03)
	r.POST("/hand04", handler.MyHandler04)

	r.Run(":9698")
}
