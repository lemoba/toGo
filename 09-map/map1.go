package main

import "fmt"

func main() {
	// 方法一：声明m1为map类型
	var m1 map[string]string
	if m1 == nil {
		fmt.Println("m1 is nil")
	}

	m1 = make(map[string]string, 10)
	m1["one"] = "golang"
	m1["two"] = "php"
	m1["three"] = "python"

	fmt.Println(m1)

	// 方法二：
	m2 := make(map[int]string)
	m2[1] = "java"
	m2[2] = "c++"
	m2[3] = "rust"

	fmt.Println(m2)

	// 方法三：
	m3 := map[string]string{
		"one":   "php",
		"two":   "python",
		"three": "erlang",
	}

	fmt.Println(m3)
}
