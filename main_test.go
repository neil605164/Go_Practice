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

func TestAdd(t *testing.T) {
	a := add(10, 20)
	if a != 30 {
		t.Log("It's Error")
	} else {
		t.Log("It's True")
	}
}
