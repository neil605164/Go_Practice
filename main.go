package main

import "fmt"

type age interface  {
	int32 | float32
}

type Data[T age] struct {
	data []T
}

func (d *Data[T]) addData(value ...T) {
	for _,v := range value{
		d.data = append(d.data, v)
	}
}

func (d *Data[T]) sum() (s T) {
	for _,item := range d.data {
		s+=item
	}

	return
}

func main() {
	data1 := []int32{10, 20, 30, 40, 50}
    data2 := []float32{10.1, 20.2, 30.3, 40.4, 50.5}

	d1 := Data[int32]{}
    d2 := Data[float32]{}

	// set value
	d1.addData(data1...)
    d2.addData(data2...)

	// calculate
	sum1 := d1.sum()
    sum2 := d2.sum()

	fmt.Printf("sum1: %v (%T)\n", sum1, sum1)
    fmt.Printf("sum2: %v (%T)\n", sum2, sum2)
}
