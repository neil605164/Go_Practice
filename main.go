package main

import "fmt"

func main() {

	fmt.Println("平均值為：", average([]float64{1, 3, 5, 7}))
	fmt.Println("A + B = ", add(10, 20))
	// fmt.Println("Hello World")
}

func average(number []float64) float64 {
	total := 0.0

	for _, x := range number {
		total += x
	}

	return total / float64(len(number))
}

func add(a, b int) int {
	return a + b
}
