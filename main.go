package main

import "fmt"

func main() {

	fmt.Println(average([]float64{1, 3, 5, 7}))
	// fmt.Println("Hello World")
}

func average(number []float64) float64 {
	total := 0.0

	for _, x := range number {
		total += x
	}

	return total / float64(len(number))
}
