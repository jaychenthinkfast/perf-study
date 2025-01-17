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
	img := createColWorkersBuffered(width, height)
	// 保存图片到文件
	if err = png.Encode(ff, img); err != nil {
		log.Fatal(err)
	}
}

func createColWorkersBuffered(width, height int) image.Image {
	m := image.NewGray(image.Rect(0, 0, width, height))

	c := make(chan int, width) // 缓冲区容量为width，只需要入队width个任务

	var w sync.WaitGroup
	for n := 0; n < numWorkers; n++ {
		w.Add(1)
		go func() {
			for i := range c {
				// 循环处理每一个列的数据
				for j := 0; j < height; j++ {
					m.Set(i, j, pixel(i, j, width, height))
				}
			}
			w.Done()
		}()
	}

	// 只需要将每一列入队
	for i := 0; i < width; i++ {
		c <- i
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
