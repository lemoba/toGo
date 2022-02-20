package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (h *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human // SuperMan继承了Human类的方法

	level int
}

// 重定义父类的方法eat()
func (h *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

// 子类的新方法
func (h *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func main() {
	h := Human{"lisi", "female"}
	h.Eat()
	h.Walk()

	// 定义一个子类对象
	// s := SuperMan{Human{"lisi", "female"}, 88}
	var s SuperMan
	s.name = "lisi"
	s.sex = "female"
	s.level = 88

	s.Eat()  // 子类的方法
	s.Fly()  // 子类的方法
	s.Walk() // 父类的方法
}
