package main

import "fmt"

func main() {
	var a string

	// pair <type:string, value:"ranen">
	a = "ranen"

	// pair <type:string, value:"ranen">
	var allType interface{}
	allType = a

	str, _ := allType.(string)
	fmt.Println(str)
}
