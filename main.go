package main

import (
	"Go_Practice/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 載入 router
	router.LoadApiRouter(r)

	r.Run(":9698")
}
