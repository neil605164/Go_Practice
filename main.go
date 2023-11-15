/*
	Design Pattern: Method Chaining
*/

package main

import (
	"Go_Practice/mychain"
)

func main() {
	c := mychain.NewMyChain()
	c.WithName("neil").WithAge(30).PrintInfo()
}
