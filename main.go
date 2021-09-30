package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	x := 1563847412

	// 如果數字為 0 直接回傳
	if x == 0 {
		return
	}

	isNegative := false

	// 檢查是否為正整數
	if x < 0 {
		x = -x
		isNegative = true
	}

	// int to string
	strNum := strconv.Itoa(x)

	// 拆解 string 將其丟入 array list
	arrList := []string{}
	for _, v := range strNum {
		arrList = append(arrList, string(v))
	}

	// 組合資料
	strRes := ""

	if isNegative {
		strRes = "-"
	}

	for i := len(arrList) - 1; i >= 0; i-- {
		strRes += arrList[i]
	}

	res, err := strconv.Atoi(strRes)
	if err != nil {
		log.Printf("Error Message is %v", err.Error())
	}

	// 檢查資料長度
	if res > int(math.Pow(2, 31))-1 || res < int(math.Pow(-2, 31)) {
		fmt.Println(0)
		return
	}

	fmt.Println(res)
}
