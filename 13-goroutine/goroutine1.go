package main

import (
	"fmt"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	// 创建一个goroutine 去执行newTask()
	go newTask()

	i := 0
	for {
		i++
		fmt.Printf("main goroutine: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
