package main

import (
	"fmt"
)

func main() {
	// 定义一个channel
	c := make(chan int)

	go func() {
		defer fmt.Println("goroutine finished!")
		fmt.Println("goroutine running...")
		c <- 666 // 将666 发送给c
	}()

	num := <-c // 从管道c中接受数据, 并赋值给num
	// num 如果不从管道中读取c 那么上面c <- 666 就会发生阻塞

	fmt.Println("num = ", num)
	fmt.Println("main goroutine finished...")
}
