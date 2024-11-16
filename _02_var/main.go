package main

import "fmt"

// 四种定义变量的方式
var i = 1

// j := 2
func main() {
	//1. 定义单个变量

	//1.1
	//默认值是0
	var a int
	fmt.Println("a=", a)

	//1.2
	var b int = 100
	fmt.Println("b=", b)

	//1.3
	var c = 200
	fmt.Println("c=", c)

	fmt.Println(b + c)
	fmt.Println(a + b + c)

	//1.4 常用(只能定义局部变量)
	fmt.Println("---------------")
	d := 300
	fmt.Println("d=", d)
	fmt.Printf("type of d = %T\n", d) //自动推导类型

	e := 3.14159
	fmt.Printf("type of e = %T\n", e) //自动推导类型

	f := "3.1415"
	fmt.Printf("type of f = %T\n", f) //自动推导类型

	fmt.Println("---------------")
	fmt.Println(i)
	//fmt.Println(j)   //定义全局变量报错

	//2.定义多变量
	var aa, bb = 123, "123"
	fmt.Println("aa=", aa, ",bb=", bb)
	var (
		cc int  = 111
		dd      = 3.14
		ee bool = false
	)
	fmt.Println("cc=", cc)
	fmt.Println("dd=", dd)
	fmt.Println("ee=", ee)
}
