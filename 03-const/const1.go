package main

import "fmt"

// 使用const 来定义枚举
const (
	// 可以在const()添加一个关键字iota, 每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = 10 * iota // iota = 0 - 0
	SHANGHAI             // iota = 1 - 20
	SHENZHEN             // iota = 2 - 30
)

const (
	a, b = iota + 1, iota + 2 // itoa = 0 - 1, 2
	c, d                      // itoa = 1 - 2, 3
	e, f                      // itoa = 2 - 3, 4

	g, h = iota * 2, iota * 3 // itoa = 3 - 6, 9
	i, j                      // itoa = 4 - 8, 12
	k, l                      // itoa = 5 - 10, 15
)

func main() {
	// 常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	// length = 100 // 常量是不能重新赋值的

	fmt.Println(BEIJING, SHANGHAI, SHENZHEN)

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)

	// itoa只能配合const()食用
	// var aa int = itoa //undefined
}
