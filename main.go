package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10

	fmt.Println(isPalindrome(x))

}

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x > int(math.Pow(2, 31)-1) {
		return false
	}

	var intL []int
	tmp := x
	for tmp != 0 {
		intL = append(intL, tmp%10)
		tmp /= 10
	}

	time := 1
	tmp = 0
	for i := len(intL) - 1; i >= 0; i-- {
		tmp += intL[i] * time
		time *= 10
	}

	return x == tmp
}
