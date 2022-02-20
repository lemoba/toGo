package main

import "fmt"

func main() {

	// 压栈操作：先进后出
	defer fmt.Println("main end")
	defer fmt.Println("main end1")

	fmt.Println("main: go1")
	fmt.Println("main: go2")
	fmt.Println("main: go3")

	defer fmt.Println("main end2")
}
