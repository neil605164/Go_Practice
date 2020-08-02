package main

import (
	"fmt"
)

func main() {
	arrs := []int{1, 1, 2, 2, 3}
	res := countGoodTriplets(arrs, 0, 0, 1)

	fmt.Println(res)
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	sum := 0

	var i, j, k int
	for i = range arr {
		for j = len(arr) - 1; j > i; j-- {
			for k = len(arr) - 1; k > j; k-- {

				if withBranch(arr[i]-arr[j]) > a {
					continue
				}

				if withBranch(arr[j]-arr[k]) > b {
					continue
				}

				if withBranch(arr[i]-arr[k]) > c {
					continue
				}

				res := fmt.Sprintf("I=%d, J=%d, K=%d", i, j, k)
				fmt.Println(res)
				sum++
			}
		}
	}

	return sum
}

func withBranch(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
