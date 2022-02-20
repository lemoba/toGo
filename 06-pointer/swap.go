package main

import "fmt"

// func swap_value(a int, b int) {
// 	var temp int
// 	temp = a
// 	a = b
// 	b = temp
// }

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a, b := 3, 5
	fmt.Println("a = ", a, "b = ", b)
	swap(&a, &b)
	//swap_value(a, b)
	fmt.Println("a = ", a, "b = ", b)

	var p *int
	p = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	var p2 **int
	p2 = &p
	fmt.Println(&p)
	fmt.Println(p2)
	fmt.Println(**p2)
}
