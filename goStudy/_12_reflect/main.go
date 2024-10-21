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

// func (u *User) Call() {  //使用指针反射不到方法
func (u User) Call() {
	fmt.Println("User is called....")
	fmt.Printf("%v\n", u)
}

func main() {
	u := User{
		id:   123,
		name: "小陈",
		age:  24,
	}
	uType := reflect.TypeOf(u)
	uValue := reflect.ValueOf(u)
	fmt.Println("uType: ", uType)
	fmt.Println("uValue: ", uValue)
	fmt.Println(u)
	fmt.Println(reflect.TypeOf(u.id))
	fmt.Println(reflect.ValueOf(u.id))
	fmt.Println(reflect.TypeOf(uValue))
	fmt.Println("-------------------")
	//反射获取字段以及值
	for i := 0; i < uType.NumField(); i++ {
		field := uType.Field(i)
		value := uValue.Field(i)
		fmt.Println(field, "--", value)
	}
	fmt.Println("---------------------")
	//反射获取方法
	for i := 0; i < uType.NumMethod(); i++ {
		method := uType.Method(i)
		fmt.Println(method)
	}

}
