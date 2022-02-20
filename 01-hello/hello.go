package main // 程序的包名类似java的package

import (
	"fmt"
	"time"
)

// main函数
func main() { // 括号不能分行写
	time.Sleep(1 * time.Second)
	fmt.Println("hello world")
}

/*
总结
* 1. 括号不能分行写必须紧跟当前行
* 2. 末尾可加可不加分号(建议不要加分号）
* 3. 可通过import导入其他包(类似python中的import)
*/
