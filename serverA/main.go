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

func Service01(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"a": "service a",
	})
}
