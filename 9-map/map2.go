package main

import "fmt"

func printMap(cityMap map[string]string) {
	// 遍历 引用传递
	for k, v := range cityMap {
		fmt.Printf("key = %s value = %s\n", k, v)
	}
}
func main() {
	cityMap := make(map[string]string)

	// 添加
	cityMap["China"] = "Shanghai"
	cityMap["USA"] = "NewYork"
	cityMap["Japan"] = "Tokyo"

	// 遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "Japan")

	// 修改
	cityMap["USA"] = "DC"

	fmt.Printf("----------------\n")

	printMap(cityMap)
}
