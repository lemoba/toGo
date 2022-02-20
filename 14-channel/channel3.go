package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c) // 如果不关闭 下面还在读取就会造成死锁
		// close可以关闭一个channel
	}()

	for {
		// ok如果为true表示channel还没有关闭，如果false表示channel已经关闭
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("Main Finished...")
}
