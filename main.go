package main

import (
	"fmt"
	"log"
	"strconv"
)

// 數字轉字串
func print01(num int) string {
	return fmt.Sprintf("%d", num)
}

// 數字轉字串
func print02(num int64) string {
	return strconv.FormatInt(num, 10)
}

// 數字轉字串
func print03(num int) string {
	return strconv.Itoa(num)
}

func main() {
	log.Println(print01(100))
	log.Println(print02(100))
	log.Println(print03(100))
}
