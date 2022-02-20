package main

import "fmt"

func main() {
	var nums = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)

	nums = append(nums, 4) //追加元素
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)

	nums = append(nums, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)

	nums = append(nums, 6) // 当长度小于1024时 扩充为原来的2倍, 否则扩充为原来的1.25倍
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)

	var s3 = make([]int, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s3), cap(s3), s3)
	s3 = append(s3, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s3), cap(s3), s3)
	//fmt.Println(nums)
}
