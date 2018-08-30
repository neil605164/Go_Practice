package main

import (
	"fmt"
	"time"
)

func main() {
	msg := "Let's Go"
	// 印出結果為："Let's GOGOGO"
	go func() {
		fmt.Println("msg:", msg)
	}()

	// 印出結果為："Let's GO"
	go func(input string) {
		fmt.Println("input:", input)
	}(msg)

	msg = "Let's GOGOGO"
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World")
}
