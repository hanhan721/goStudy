package main

import (
	"fmt"
	"reflect"
)

type User struct {
	id   int16
	name string
	age  int
}

func main() {
	u := User{
		id:   123,
		name: "小陈",
		age:  24,
	}
	uType := reflect.TypeOf(u)
	uValue := reflect.ValueOf(u)
	fmt.Println(uType)
	fmt.Println(uValue)
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(uValue))

}
