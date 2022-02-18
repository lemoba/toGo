package main

import "fmt"

func prinArray(arr [4]int) {
	for index, value := range arr {
		fmt.Println("index = ", index, " value = ", value)
	}
}

func main() {
	// 固定长度的数组 默认值为0
	var arr1 [10]int

	arr2 := [10]int{1, 2, 3}

	arr3 := [4]int{11, 22, 33, 44}

	for i := 0; i < len(arr1); i++ {
		fmt.Printf("%d ", arr1[i])
	}

	//pinArray(arr2) // 类型不匹配 cannot use arr2 (type [10]int) as type [4]int
	prinArray(arr3) // 值传递

	fmt.Println()

	// 查看数组的数据类型
	fmt.Printf("arr1 type = %T\n", arr1)
	fmt.Printf("arr2 type = %T\n", arr2)
	fmt.Printf("arr3 type = %T\n", arr3)
}
