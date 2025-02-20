package main

import "testing"

type Item struct {
	id  int
	val [2048]byte
}

// 用下标遍历[]struct{}
func BenchmarkSliceStructByFor(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := 0; j < len(items); j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}

// 用for range下标遍历[]struct{}
func BenchmarkSliceStructByRangeIndex(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 用for range值遍历[]struct{}的元素
func BenchmarkSliceStructByRangeValue(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
