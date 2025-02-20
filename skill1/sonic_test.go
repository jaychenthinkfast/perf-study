package main

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type UserTwo struct {
	Name string
	Age  int
}

// sonic序列化/反序列化
func TestSonic(t *testing.T) {
	user := UserTwo{Name: "test", Age: 18}
	// sonic序列化
	data, err := sonic.Marshal(user)
	if err != nil {
		panic(err)
	}
	var newUser UserTwo
	err = sonic.Unmarshal(data, &newUser)
	if err != nil {
		panic(err)
	}
}

// sonic 序列化基准测试
func BenchmarkSonicMarshal(b *testing.B) {
	user := UserTwo{Name: "test", Age: 18}
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// sonic 反序列化基准测试
func BenchmarkSonicUnmarshal(b *testing.B) {
	user := UserTwo{Name: "test", Age: 18}
	data, err := sonic.Marshal(user)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var newUser UserTwo
		err = sonic.Unmarshal(data, &newUser)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// 标准库序列化基准测试
func BenchmarkStdMarshal(b *testing.B) {
	user := UserTwo{Name: "test", Age: 18}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(user)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// 标准库反序列化基准测试
func BenchmarkStdUnmarshal(b *testing.B) {
	user := UserTwo{Name: "test", Age: 18}
	data, err := json.Marshal(user)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var newUser UserTwo
		err = json.Unmarshal(data, &newUser)
		if err != nil {
			b.Fatal(err)
		}
	}
}
