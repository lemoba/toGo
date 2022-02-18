package main

import "fmt"

func main() {
	s := []int{1, 2, 3}

	//[0, 1]
	//s1 := s[0:3] // 前闭后开 s1指向s底层数组

	// s1[0] = 100 // 原slice也会被修改

	//copy 可以将底层数组的slice一起拷贝
	s2 := make([]int, 3) // [0, 0, 0]
	copy(s2, s)
	s2[0] = 100

	fmt.Println(s2)
	fmt.Println(s)
}
