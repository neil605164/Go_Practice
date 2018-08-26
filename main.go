package main

import "fmt"

func checkValue(num int) {
	switch num {
	case 0:
		fallthrough // 若符合，繼續往下跑
	case 1:
		fmt.Printf("num is %d", num)
		fmt.Println()
	case 2, 3:
		fmt.Println("值為：", num)
	}
}

func main() {
	checkValue(0)
	checkValue(1)
	checkValue(2)
	checkValue(3)
}
