package main

import (
	// "github.com/lemoba/toGo/5-init/lib1"
	_ "github.com/lemoba/toGo/5-init/lib1" // 匿名别名
	mylib "github.com/lemoba/toGo/5-init/lib2"
	//. "github.com/lemoba/toGo/5-init/lib2" // 导入包中所有方法,不建议使用
)

func main() {
	// lib1.Lib1Test()
	mylib.Lib2Test()
	// Lib2Test()
}
