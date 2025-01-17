package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"runtime/trace"
	"sync"
)

const (
	output     = "out.png"
	width      = 2048
	height     = 2048
	numWorkers = 8
)

func main() {
	// 开启trace输出到stdout上
	trace.Start(os.Stdout)
	defer trace.Stop()

	// 创建待生成的图片文件
	ff, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	// 生成曼德勃罗特图
	img := createWorkersBuffered(width, height)
	// 保存图片到文件
	if err = png.Encode(ff, img); err != nil {
		log.Fatal(err)
	}
}

// createWorkers 并行优化
func createWorkersBuffered(width, height int) image.Image {
	// 创建image对象
	m := image.NewGray(image.Rect(0, 0, width, height))
	type px struct{ x, y int }
	c := make(chan px, width*height) // 修改这里，缓冲区大小为width*height
	var w sync.WaitGroup
	// numWorkers为8，创建8个工作者
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			// 从channel出队像素对象并处理
			for px := range c {
				m.Set(px.x, px.y, pixel(px.x, px.y, width, height))
			}
			w.Done()
		}()
	}
	// 遍历每一个像素点
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			// 向channel中入队像素对象
			c <- px{i, j}
		}
	}
	close(c)
	w.Wait()
	return m
}

// pixel returns the color of a Mandelbrot fractal at the given point.
func pixel(i, j, width, height int) color.Color {
	// Play with this constant to increase the complexity of the fractal.
	// In the justforfunc.com video this was set to 4.
	const complexity = 1024

	xi := norm(i, width, -1.0, 2)
	yi := norm(j, height, -1, 1)

	const maxI = 1000
	x, y := 0., 0.

	for i := 0; (x*x+y*y < complexity) && i < maxI; i++ {
		x, y = x*x-y*y+xi, 2*x*y+yi
	}

	return color.Gray{uint8(x)}
}

func norm(x, total int, min, max float64) float64 {
	return (max-min)*float64(x)/float64(total) - max
}
