package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

func zero(xPtr *int) {
	*xPtr = 0
}
