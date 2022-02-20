package main

import "fmt"

// 类名首字母大写，表示其他包也能够访问
type Hero struct {
	// 类属性首字母大写，表示该属性可以对外访问(public)否则只能在类内部使用(private)
	Name  string
	Ad    int
	Level int
}

// 方法接受者类似this或self
// func (h Hero) GetName() string {
// 	return h.Name
// }

// func (h *Hero) SetName(name string) {
// 	h.Name = name
// }

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) SetName(name string) {
	h.Name = name
}

func main() {
	// 创建一个对象
	hero := Hero{
		Name:  "ranen",
		Ad:    111,
		Level: 1,
	}
	name := hero.GetName()
	fmt.Println("name = ", name)
	hero.SetName("king")
	name1 := hero.GetName()
	fmt.Println("name = ", name1)
}
