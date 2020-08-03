package main

func main() {
	nums1 := []int{4, 0, 0, 0, 0, 0}
	nums2 := []int{1, 2, 3, 5, 6}
	merge(nums1, 1, nums2, 5)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	var i, j, k int
	k = (m + n) - 1

	// 比對 num1 與 num2，數字大者一順序塞入 nums1 後方
	// 直到 i,j 某一條件不成立
	for i, j = m-1, n-1; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for i >= 0 {
		nums1[k] = nums1[i]
		i--
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}
