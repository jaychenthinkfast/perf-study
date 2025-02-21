package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange(安排）
	a := 5
	b := 3
	expected := 8

	// Act（行动）
	result := Add(a, b)

	// Assert（断言）
	assert.Equal(t, expected, result)
}

// 通过定义结构体切片来构建测试数据表格，接着在测试函数中遍历这个表格，对每组数据执行相同的测试逻辑，从而实现 table-driven 模式。
func TestAdd0(t *testing.T) {
	tests := []struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}{
		{"Add positive numbers", 1, 2, 3},
		{"Add negative numbers", -1, 1, 0},
		{"Add zero", 0, 0, 0},
		// 更多测试用例...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.inputA, tt.inputB)
			assert.Equal(t, tt.expected, result, "结果不符合期望")
		})
	}
}

func TestAddWithConvey(t *testing.T) {
	Convey("关于Add函数的测试", t, func() {
		Convey("正常情况的测试", func() {
			Convey("两个正数相加", func() {
				result := Add(2, 3)
				So(result, ShouldEqual, 5)
			})
			Convey("一个正数和一个负数相加", func() {
				result := Add(5, -3)
				So(result, ShouldEqual, 2)
			})
		})
		Convey("边界情况的测试", func() {
			Convey("两个零相加", func() {
				result := Add(0, 0)
				So(result, ShouldEqual, 0)
			})
			Convey("一个数与最大整数相加", func() {
				result := Add(int(math.MaxInt32), 1)
				So(result, ShouldEqual, int(math.MaxInt32)+1)
			})
		})
	})
}
