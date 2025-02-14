package main

import (
	"testing"
)

// 泛型实现的最大值函数benchmark
func BenchmarkGenerics(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(1, 2)
	}
}

// MaxInt函数benchmark
func BenchmarkRegular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxInt(1, 2)
	}
}
