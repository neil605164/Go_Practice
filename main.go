// Struct 教學

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/", getData)
	r.Run()
}

type jsonData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	NotUse  int         `json:"-"`
}

type obj struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getData(c *gin.Context) {
	jsonData := jsonData{
		Status:  "OK",
		Message: "Access",
		Data: []obj{
			obj{Name: "Neil", Age: 25},
			{Name: "Neil1", Age: 25},
			{Name: "Neil2", Age: 25},
		},
		NotUse: 1,
	}
	c.JSON(200, jsonData)
}
