package main

import "fmt"

func main() {
	nums := []int{9}
	// nums := []int{8, 9}
	// nums := []int{9, 9}
	// nums := []int{8, 9, 9}
	// nums := []int{1, 2, 3}
	// nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res := plusOne(nums)
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	carry := true
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
				if i == 0 {
					digits = append(digits, 0)
					digits[0] = 1
				}
				carry = true
			} else {
				digits[i] += 1
				carry = false
			}
		} else {
			return digits
		}
	}

	return digits
}
