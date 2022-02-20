package main

import "fmt"

// 声明一种新的数据类型myint, 是int的别名
type myint int

// 定义一个结构体
type Book struct {
	title  string
	author string
}

func changeBook(book Book) {
	//值传递
	book.title = "aaaa"
}

func changeBook2(book *Book) {
	// 指针传递
	book.title = "bbbb"
}

func main() {
	// var a myint = 10
	// fmt.Println("a = ", a)
	// fmt.Printf("type of a = %T\n", a)
	var book1 Book
	book1.title = "Go"
	book1.author = "ranen"
	fmt.Printf("%v\n", book1)
	changeBook(book1)
	fmt.Printf("%v\n", book1)
	changeBook2(&book1)
	fmt.Printf("%v\n", book1)
}
