package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	a := "aaa"
	b := "bbb"
	_, err := Max(a, b)
	if err != nil {
		panic(err)
	}
}

// MaxInt函数benchmark
func BenchmarkRegular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MaxInt(1, 2)
	}
}

// 反射实现的最大值函数benchmark
func BenchmarkReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Max(1, 2)
	}
}
