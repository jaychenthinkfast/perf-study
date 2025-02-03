package main

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

// Test1kGCDuration 测试小规模数据gc时长
func Test1kGCDuration(t *testing.T) {
	size := 1000
	m := GenerateStringMap(size)
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m["1"]
}

// 测试大规模数据gc时长
func Test500wGCDuration(t *testing.T) {
	size := 5000000
	m := GenerateStringMap(size)
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m["1"]
}
func GenerateStringMap(size int) map[string]string {
	// 在这里执行一些可能会触发GC的操作，例如创建大量对象等
	// 以下示例创建一个较大的map并填充数据
	m := make(map[string]string)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("val_%d", i)
		m[key] = value

	}
	return m
}

func timeGC() time.Duration {
	// 记录GC开始时间
	gcStartTime := time.Now()
	// 手动触发GC，以便更准确地测量此次操作相关的GC时长
	runtime.GC()

	// 计算总的GC时长
	gcCost := time.Since(gcStartTime)
	return gcCost
}

func TestSmallStruct(t *testing.T) {
	type SmallStruct struct {
		data [128]byte
	}
	m := make(map[int]SmallStruct)
	size := 5000000
	for i := 0; i < size; i++ {
		m[i] = SmallStruct{}
	}
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m[1]
}
func TestBigStruct(t *testing.T) {
	type BigStruct struct {
		data [129]byte
	}
	m := make(map[int]BigStruct)
	size := 5000000
	for i := 0; i < size; i++ {
		m[i] = BigStruct{}
	}
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m[1]
}
