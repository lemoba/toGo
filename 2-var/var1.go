package main

import "fmt"

// 声明全局变量 方法一、二、三是可以的
var gA int = 100
var gB = 200

//方法4
//syntax error: non-declaration statement outside function body
// gC := 300

func main() {
	// 方法一：声明一个变量 默认值为0
	var a int
	fmt.Printf("type of a = %T\n", a)
	fmt.Println("a = ", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 1
	fmt.Printf("type of b = %T\n", b)
	fmt.Println("b = ", b)

	// 方法三：在初始化的时候，可以省去数据类型，通过自动匹配当前的变量的数据类型
	var c = 100
	fmt.Printf("type of c = %T\n", c)
	fmt.Println("c = ", c)

	// 方法四：(常用的方法)省去var关键字，直接自动匹配
	d := 100
	fmt.Printf("type of d = %T\t value = %d\n", d, d)

	e := "hello world"
	fmt.Printf("type of e = %T e = %s\n", e, e)

	f := 100.11
	fmt.Printf("type of f = %T f = %f\n", f, f)

	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)
	// fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy int = 100, 200
	z1, z2 := 11, 22
	var kk, ll = 55, "hello"
	fmt.Println("xx = ", xx, ", yy = ", yy)
	fmt.Println("z1 = ", z1, ", z1 = ", z2)
	fmt.Println("kk = ", kk, ", ll = ", ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}

/*
总结
* 1. 声明全局变量使用方法一、二、三
* 2. 方法四只能在函数体内声明
*/
