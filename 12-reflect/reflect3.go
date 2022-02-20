package main

import (
	"fmt"
	"reflect"
)

type resume struct {
	Name string `info:"name" doc:"ranen"`
	Sex  string `info:"sex"`
}

func findTag(str interface{}) {
	t := reflect.TypeOf(str).Elem()
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Printf("info = %s doc = %s\n", tagInfo, tagDoc)
	}
}

func main() {
	var re resume
	findTag(&re)
}
