package main

import "fmt"

//本质是一个指针
type Animal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// 具体的类
type Cat struct {
	color string //猫的颜色
}

// 实现接口
func (c *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

// 具体类
type Dog struct {
	color string
}

// 实现接口
func (c *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}

func (c *Dog) GetColor() string {
	return c.color
}

func (c *Dog) GetType() string {
	return "Dog"
}

func showAnimal(animal Animal) {
	animal.Sleep() // 多态
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

func main() {
	/*
		var animal Animal // 接口的数据类型， 父类指针

		animal = &Cat{"Green"}
		animal.Sleep() // 调用Cat的Sleep()方法

		animal = &Dog{"Yellow"}
		animal.Sleep() // 调用Cat的Sleep()方法
	*/

	cat := Cat{"Green"}
	showAnimal(&cat)

	dog := Dog{"Yellow"}
	showAnimal(&dog)
}
