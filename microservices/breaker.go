package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromicro/go-zero/core/breaker"
)

type mockError struct {
	status int
}

func (e mockError) Error() string {
	return fmt.Sprintf("HTTP STATUS: %d", e.status)
}

func main() {
	for i := 0; i < 1000; i++ {
		if err := breaker.Do("test", mockRequest); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func mockRequest() error {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	num := r.Intn(100)
	if num%4 == 0 {
		return nil
	} else if num%5 == 0 {
		return mockError{status: 500}
	}
	return fmt.Errorf("dummy")
}
