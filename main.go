package main

import "fmt"

// 執行main前，會先執行init
func init() {
	fmt.Println("Hello Init")
}

func main() {
	fmt.Println("Hello World")
}
