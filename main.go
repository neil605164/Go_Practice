package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(
		// 比對字串是否有符合條件
		strings.Contains("test", "es"),

		// 比較字串是否相符合
		strings.Compare("A", "a"),

		// 比較任何符號
		strings.ContainsAny("I'm a boy", "'"),

		// 字串相加
		strings.Join([]string{"a", "b", "c"}, "-"),

		// 重複印出指定字串
		strings.Repeat("#", 5),

		// 字串轉換成大寫
		strings.ToLower("Fdwgb"),

		// 字串轉換成小寫
		strings.ToUpper("FredwC"),

		// 計算某字串出現幾次
		strings.Count("no no no no yes", "no"),
	)
}
