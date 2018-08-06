package main

import (
	"fmt"
)

func main() {
	s := square{width: 3, height: 4}
	// c := circle{radius: 5}

	measure(s)
	// measure(c)
}

// 定義 interface: grometry，且必須實現以下func
// 若須使用grometry這一個interface，帶入的『參數』必須實現有area，perimeter這兩個func
type grometry interface {
	// func名稱(變數, 帶入的資料型態) (回傳的變數, 回傳的資料型態)
	area() float64
	perimeter() float64
}

//  定義方形的規定
// 可視為宣告一個class，底下有兩個func(area,perimeter)
type square struct {
	width, height float64
}

// // 定義圓形的規定
// // 可視為宣告一個class，底下有兩個func(area,perimeter)
// type circle struct {
// 	radius float64
// }

// 方形面積
func (s square) area() float64 {
	return s.width * s.height
}

// 方形週長
func (s square) perimeter() float64 {
	return 2*s.width + 2*s.height
}

// // 圓形面積
// func (c circle) area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// // 圓週率
// func (c circle) perimeter() float64 {
// 	return 2 * math.Pi * c.radius
// }

// 不太懂
func measure(g grometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
