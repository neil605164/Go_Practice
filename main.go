package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 26

	res := twoSum(nums, target)
	fmt.Println(res)
}

func twoSum(nums []int, target int) []int {

	m := make(map[int]int)
	for i, n := range nums {
		if j, ok := m[target-n]; ok {
			return []int{j, i}
		}
		m[n] = i
	}
	return []int{}
}
