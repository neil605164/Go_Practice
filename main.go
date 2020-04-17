package main

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {

		p := &sync.Pool{
			New: func() interface{} {
				return 0
			},
		}

		a := p.Get().(int)

		fmt.Println("A ==>", a) // 0
		p.Put(1)
		runtime.GC()

		// time.Sleep(time.Second * 1)
		b := p.Get().(int)
		fmt.Println("A ==>", a) // 0
		fmt.Println("B ==>", b) // 1

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8787") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
