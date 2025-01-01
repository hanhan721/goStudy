package main

import "fmt"

// 定义结构体,类似java的实体类
type Book struct {
	title string
	auth  string
	price float32
}

func changBook(book *Book) {
	book.price = 188.88
}

func main() {
	book := Book{"三体", "刘慈欣", 68.88}
	fmt.Printf("type of book = %T\n", book)
	fmt.Println(book)
	changBook(&book) //得使用指针才能修改结构体内的值
	fmt.Println(book)

	fmt.Println("----------------------------------")

	student := Student{}
	student.SetName("张三")
	student.SetAge(14)
	fmt.Println(&student)
	fmt.Println(student.GetName())

	student1 := new(Student) //指针类型
	//student1 := &Student{} 和上边一样
	student1.SetAge(15).SetName("李四")
	fmt.Println(student1)
	fmt.Printf("student类型: %T\n", student)
	fmt.Printf("student1类型: %T\n", student1)
}
