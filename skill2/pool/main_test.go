package pool

import "testing"

func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 请求处理逻辑
		WriteBufferNoPool()
	}
}

func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 请求处理逻辑
		WriteBufferWithPool()
	}
}
