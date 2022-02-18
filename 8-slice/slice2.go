package main

import "fmt"

func main() {

	//s1 := []int{1, 2, 3, 4} //方法1 声明切片，并初始化 len = 4

	//var s2 []int // 方法2 声明s2为slice,但没有分配空间
	// s2 = make([]int, 4) // 使用make开辟4个空间

	//var s2 = make([]int, 4) // 方法3 声明切片并初始化空间，初始化值为0
	s2 := make([]int, 4) // 方法3

	// s2[1] = 10

	// 判断slice是否为空
	if s2 == nil {
		fmt.Printf("s2 is no space\n")
	} else {
		fmt.Printf("s2 is space\n")
	}

	fmt.Printf("len = %d, slice = %v\n", len(s2), s2)
}
