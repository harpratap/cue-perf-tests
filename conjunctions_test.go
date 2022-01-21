package main

import (
	"testing"
)

func BenchmarkConjunctionDouble(b *testing.B) {
	for n := 0; n < b.N; n++ {
		conjunctionDouble()
	}
}

func BenchmarkConjunctionSingle(b *testing.B) {
	for n := 0; n < b.N; n++ {
		conjunctionSingle()
	}
}
