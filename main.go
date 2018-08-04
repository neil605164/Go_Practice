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

// class 概念
type jsonData struct {
	Username string      `json:"username"`
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	NotUse   int         `json:"-"`
}

type obj struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *jsonData) getMe() (myname string) {
	myname = "My name is: " + c.Username
	return
}

func getData(c *gin.Context) {
	data := jsonData{
		Username: "Neil",
		Status:   "OK",
		Message:  "Access",
		Data: []obj{
			obj{Name: "Neil", Age: 25},
			{Name: "Neil1", Age: 25},
			{Name: "Neil2", Age: 25},
		},
		NotUse: 1,
	}

	myname := data.getMe()
	c.JSON(200, myname)
}
