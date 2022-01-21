package main

import (
	"testing"
)

func BenchmarkDefault0(b *testing.B) {
	for n := 0; n < b.N; n++ {
		default0()
	}
}

func BenchmarkDefault10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		default10()
	}
}
