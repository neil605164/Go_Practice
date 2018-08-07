package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Add 时间相加
	now := time.Now()

	// 1分钟前
	m1 := now.Add(-time.Minute)
	fmt.Println(m1)

	// 8个小时前
	h1 := now.Add(-time.Hour * 8)
	fmt.Println(h1)

	// 一天前
	d1 := now.Add(-time.Hour * 24)
	fmt.Println(d1)

	printSplit(50)

	// 10分钟后
	m1 = now.Add(time.Minute)
	fmt.Println(m1)

	// 8小时后
	h1 = now.Add(time.Hour * 8)
	fmt.Println(h1)

	// 一天后
	d1 = now.Add(time.Hour * 24)
	fmt.Println(d1)

	printSplit(50)

	// Sub 计算两个时间差
	subM := now.Sub(m1)
	fmt.Println(subM.Minutes(), "分钟")

	sumH := now.Sub(h1)
	fmt.Println(sumH.Hours(), "小时")

	sumD := now.Sub(d1)
	fmt.Printf("%v 天\n", sumD.Hours()/24)

}

func printSplit(count int) {
	fmt.Println(strings.Repeat("#", count))
}
