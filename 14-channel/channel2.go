package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // 有缓冲的channel
	fmt.Printf("len(c) = %d, cap(c) = %d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("子go程结束")

		for i := 0; i < 4; i++ {
			c <- i
			fmt.Printf("子go程正在运行: 发送的元素 = %d, len(c) = %d cap(c) = %d\n", i, len(c), cap(c))
		}
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 4; i++ {
		num := <-c //从c中接受并赋值
		fmt.Println("num = ", num)
	}

	fmt.Println("main 结束")
}
