package main

import "fmt"

func printArray(arr []int) {
	// _ 表示匿名的变量
	for _, v := range arr {
		fmt.Printf("value = %d\n", v)
	}
	arr[2] = 100
}

func main() {
	// 动态数组，切片slice 引用传递
	arr1 := []int{1, 2, 3, 4}
	fmt.Println(len(arr1))
	fmt.Println(arr1)

	printArray(arr1)

	fmt.Println("----------")
	for _, v := range arr1 {
		fmt.Printf("value = %d\n", v)
	}
}
