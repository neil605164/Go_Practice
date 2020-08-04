package main

import (
	"fmt"
	"math"
)

func main() {
	a := []string{"bella", "label", "roller"}

	res := commonChars(a)
	fmt.Println(res)
}

func commonChars(A []string) []string {
	tmp := make([]int, 26)
	for i := range tmp {
		tmp[i] = math.MaxInt32
	}
	tmp2 := make([]int, 26)
	// 分批執行每個字串
	for _, str := range A {

		// 檢查該字元位於 26 字母中哪個位置
		for _, c := range []byte(str) {
			tmp2[c-byte('a')]++
		}

		// 這裡不懂
		for i := 0; i < 26; i++ {
			tmp[i] = min(tmp[i], tmp2[i])
		}

		// 從新歸零 tmp2 每個位置
		for i := 0; i < 26; i++ {
			tmp2[i] = 0
		}
	}

	var rst []string
	// 將 tmp 中出現幾次的該字母，塞入 ret 中
	for i, v := range tmp {
		for k := 0; k < v; k++ {
			rst = append(rst, string([]byte{byte(i) + byte('a')}))
		}
	}
	return rst
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
