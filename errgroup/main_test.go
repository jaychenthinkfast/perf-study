package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"sync"
	"testing"
	"time"
)

var urls = []string{
	"https://chenjie.info/",
	"https://golang.org",
	"https://google.com",
}

// 定义默认超时时间为5秒
var defaultTimeout = 5 * time.Second

func TestWaitGroup(t *testing.T) {
	// 创建WaitGroup对象
	wg := sync.WaitGroup{}
	results := make([]string, len(urls))
	for index, url := range urls {
		url := url
		index := index
		// 在创建协程执行任务之前，调用Add方法
		wg.Add(1)
		go func() {
			// 任务完成后，调用Done方法
			defer wg.Done()
			// Fetch the URL with timeout.
			client := &http.Client{Timeout: defaultTimeout}
			resp, err := client.Get(url)
			if err != nil {
				return
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return
			}
			results[index] = string(body)

		}()
	}
	// 主协程阻塞，等待所有的任务执行完成
	wg.Wait()
}

func TestErrHandle(t *testing.T) {
	results := make([]string, len(urls))
	// 创建Group类型
	g := new(errgroup.Group)
	for index, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url
		index := index
		// 调用Go方法
		g.Go(func() error {
			// Fetch the URL with timeout.
			client := &http.Client{Timeout: defaultTimeout}
			resp, err := client.Get(url)
			if err != nil {
				return err // 返回错误
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return err // 返回错误
			}
			results[index] = string(body)
			return nil
		})
	}
	// Wait for all HTTP fetches to complete.
	// 等待所有任务执行完成，并对错误进行处理
	if err := g.Wait(); err != nil {
		fmt.Println("Failured fetched all URLs.")
	}
}

func TestCancel(t *testing.T) {
	results := make([]string, len(urls))
	// 用WithContext函数创建Group对象
	eg, ctx := errgroup.WithContext(context.Background())
	for index, url := range urls {
		url := url
		index := index
		// 调用Go方法
		eg.Go(func() error {
			select {
			case <-ctx.Done(): // select-done模式取消运行
				return errors.New("task is cancelled")
			default:
				// Fetch the URL with timeout.
				client := &http.Client{Timeout: defaultTimeout}
				resp, err := client.Get(url)
				if err != nil {
					return err // 返回错误
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return err // 返回错误
				}
				results[index] = string(body)
				return nil
			}
		})
	}
	// Wait for all HTTP fetches to complete.
	// 等待所有任务执行完成，并对错误进行处理
	if err := eg.Wait(); err != nil {
		fmt.Println("Failured fetched all URLs.")
	}
}

func TestLimitGNum(t *testing.T) {
	results := make([]string, len(urls))
	// 用WithContext函数创建Group对象
	eg, ctx := errgroup.WithContext(context.Background())
	// 调用SetLimit方法，设置可同时运行的最大协程数
	eg.SetLimit(2)
	for index, url := range urls {
		url := url
		index := index
		// 调用Go方法
		eg.Go(func() error {
			select {
			case <-ctx.Done(): // select-done模式取消运行
				return errors.New("task is cancelled")
			default:
				// Fetch the URL with timeout.
				client := &http.Client{Timeout: defaultTimeout}
				resp, err := client.Get(url)
				if err != nil {
					return err // 返回错误
				}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					return err // 返回错误
				}
				results[index] = string(body)
				return nil
			}
		})
	}
	// Wait for all HTTP fetches to complete.
	// 等待所有任务执行完成，并对错误进行处理
	if err := eg.Wait(); err != nil {
		fmt.Println("Failured fetched all URLs.")
	}
}
