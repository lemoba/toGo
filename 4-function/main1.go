package main

import "fmt"

func foo(a string, b int) int {
	fmt.Printf("--------foo--------\n")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

// 返回多个匿名返回值
func foo2(a string, b int) (int, int) {
	fmt.Printf("--------foo2--------\n")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 520, 13
}

// 返回多个有形参的返回值
func foo3(a string, b int) (r1, r2 int) {
	fmt.Printf("--------foo3--------\n")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	// r1 r2属于foo3的形参, 初始化默认的值为0
	// r1 r2 作用域空间 是foo3整个函数{}空间

	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	// 给有名称的返回值变量赋值
	r1 = 100
	r2 = 200

	return
}

func main() {
	r := foo("hello world", 10)
	fmt.Println("r = ", r)

	r, r2 := foo2("Love You", 20)
	fmt.Println("r = ", r)
	fmt.Println("r2 = ", r2)

	r3, r4 := foo3("Love world", 222)
	fmt.Println("r3 = ", r3)
	fmt.Println("r4 = ", r4)
}
