package main

import "fmt"

func f1() {
	fmt.Println("A")
}

func f2() {
	fmt.Println("B")
}

func f3() {
	fmt.Println("C")
}

func main() {
	defer f1()
	defer f2()
	defer f3()
}
