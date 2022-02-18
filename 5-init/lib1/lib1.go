package lib1

import "fmt"

// 函数名大写表示可对外导出，小写则表示只能在当前包中使用
func Lib1Test() {
	fmt.Println("lib1Test()...")
}
func init() {
	fmt.Println("lib1.init()...")
}
