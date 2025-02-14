package main

import "fmt"

// Max使用泛型来比较两个同类型的值（要求类型是可比较的），并返回较大的值
func Max[T int | float32](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//类型隐式传入，由编译器自动推导
	fmt.Println(Max(1, 2))
	//类型显式传入
	fmt.Println(Max[float32](1.10, 2.2))
}
