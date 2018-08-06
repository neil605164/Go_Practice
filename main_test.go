package main

import "testing"

func TestAverage(t *testing.T) {
	v := average([]float64{1, 3})
	if v != 2.0 {
		t.Log("The result is Not true")
	} else {
		t.Log("The result is true")
	}
}
