package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	mathtest "github.com/neil605164/go_pkg"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		floatList := []float64{2.2, 3.2, 5.5}

		fmt.Println(floatList)

		result := mathtest.Average(floatList)
		fmt.Println(result)

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":1212")
}
