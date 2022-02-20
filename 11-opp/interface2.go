package main

import "fmt"

//interface{}是万能数据类型
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	// interface{} 如何区分所引用的底层数据类型

	// interface{} 提供”类型断言“机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("no string")
	} else {
		fmt.Println("arg is string type, value = ", value)
		fmt.Printf("value type is %T\n", value)
	}
}

type Book struct {
	author string
}

func main() {
	book := Book{"ranen"}
	myFunc(book)
	myFunc("aaa")
	myFunc(12)
	myFunc(1.12)
}
