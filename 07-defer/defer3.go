package main

import "fmt"

/*
defer和return谁先谁后
defer先进行压栈操作然后等当前函数执行完毕后再执行
*/
func derferFunc() int {
	fmt.Println("derfer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}
func deferAndReturn() int {
	defer derferFunc()
	return returnFunc()
}

func main() {
	deferAndReturn()
}
