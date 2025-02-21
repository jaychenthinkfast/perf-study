package main

import (
	"fmt"

	"golang.org/x/sync/singleflight"
)

func main() {
	g := new(singleflight.Group)

	block := make(chan struct{})
	res1c := g.DoChan("key1", func() (interface{}, error) {
		<-block
		return "func_1", nil
	})
	res2c := g.DoChan("key1", func() (interface{}, error) {
		<-block
		return "func_2", nil
	})
	close(block)

	res1 := <-res1c
	res2 := <-res2c

	// Only the first function is executed: it is registered and started with "key",
	// and doesn't complete before the second function is registered with a duplicate key.
	fmt.Println("Equal results:", res1.Val.(string) == res2.Val.(string))
	fmt.Println("Result:", res1.Val)
}
