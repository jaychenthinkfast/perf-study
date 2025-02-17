package main

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"testing"
)

func TestMock(t *testing.T) {
	// 使用gomonkey mock函数httpGetRequest的返回
	mockData := []byte(`{"name":"killianxu","age":32}`)
	patch := gomonkey.ApplyFunc(httpGetRequest, func(url string) ([]byte, error) {
		return mockData, nil
	})
	defer patch.Reset()

	// 底层httpGetRequest的函数调用返回，会被mock
	mockUserInfo, _ := fetchUserInfo("123")

	fmt.Printf("mocked user info: %s\n", mockUserInfo)
}
