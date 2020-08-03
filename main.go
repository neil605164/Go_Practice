package main

import "fmt"

func main() {
	// nums := []int{9}
	// nums := []int{8, 9}
	// nums := []int{9, 9}
	// nums := []int{8, 9, 9}
	// nums := []int{1, 2, 3}
	nums := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	res := plusOne(nums)
	fmt.Println(res)
}

func plusOne(digits []int) []int {
	l := len(digits)
	var tmp, res []int

	mustPlus := false
	for i := l - 1; i >= 0; i-- {
		// 如果是最後一格
		if i == l-1 {
			if digits[i]+1 >= 10 {
				tmp = append(tmp, 0)

				// 長度只有 1 才需要
				if l == 1 {
					tmp = append(tmp, 1)
				}

				mustPlus = true
			} else {
				tmp = append(tmp, digits[i]+1)
				mustPlus = false
			}
		} else {
			if mustPlus {
				if digits[i]+1 >= 10 {
					tmp = append(tmp, 0)

					if i == 0 {
						tmp = append(tmp, 1)
					}

					mustPlus = true
				} else {
					tmp = append(tmp, digits[i]+1)
					mustPlus = false
				}
			} else {
				tmp = append(tmp, digits[i])
				mustPlus = false
			}
		}
	}

	// 顛倒順序
	for j := len(tmp) - 1; j >= 0; j-- {
		res = append(res, tmp[j])
	}

	return res
}
