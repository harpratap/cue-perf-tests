package main

import (
	"testing"
)

func BenchmarkDefinitionOpen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		definitionOpen()
	}
}

func BenchmarkDefinitionClosed(b *testing.B) {
	for n := 0; n < b.N; n++ {
		defintionClosed()
	}
}
