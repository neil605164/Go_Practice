package main

import "fmt"

func main() {
	input := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}

	res := missingNumber(input)
	fmt.Println(res)
}

func missingNumber(nums []int) int {
	// 先排序
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				tmp := nums[j]
				nums[j] = nums[i]
				nums[i] = tmp
			}
		}
	}

	for k, v := range nums {
		if k != v {
			return k
		}
	}

	if len(nums) != nums[len(nums)-1] {
		return len(nums)
	}

	return 0
}
