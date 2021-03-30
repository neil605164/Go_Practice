package main

import (
	"testing"
)

func BenchmarkJsoniterUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsoniterUnmarshal()
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonUnmarshal()
	}
}

func BenchmarkGJsonInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GJsonWithInterface()
	}
}

func BenchmarkGJsonWithFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GJsonWithFor()
	}
}
