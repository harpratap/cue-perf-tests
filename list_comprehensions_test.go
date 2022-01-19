package main

import (
	"testing"
)

func BenchmarkListComprehensionLong10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		listComprehensionLong10()
	}
}

func BenchmarkListComprehensionLong1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		listComprehensionLong1()
	}
}

func BenchmarkListComprehensionShort10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		listComprehensionShort10()
	}
}

func BenchmarkListComprehensionShort1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		listComprehensionShort1()
	}
}
