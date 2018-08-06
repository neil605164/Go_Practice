package main

import "fmt"

type myName struct {
	Username string `json:"username"`
}

func (c myName) getMe(string) (name string) {
	name = "Neil"
	return
}

func main() {
	var data interface {
		getMe(string) string
	}
	fmt.Println("data ---->", data)

	// 噴錯誤，原因是沒有getMe func
	// data = 1
	// fmt.Println("data ---->", data)

	// 噴錯誤，原因是沒有getMe func
	// data = "OK"
	// fmt.Println("data ---->", data)

	// 正確，因為myName這一個物件，有包含getMe的func
	data = myName{Username: "Hsieh"}
	fmt.Println("data ---->", data)
	fmt.Println()
}
