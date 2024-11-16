package main

import "fmt"

type Read interface {
	ReadBook()
}
type Write interface {
	WriteBook()
}
type Book struct {
}

func (b *Book) ReadBook() {
	fmt.Println("Reading book...")
}

func (b *Book) WriteBook() {
	fmt.Println("Writing book...")
}

func main() {
	// b: pair<type:Book, value:Book{}地址>
	b := &Book{}

	// r: pair<type:空,value:空>
	var r Read
	// b: pair<type:Book, value:Book{}地址>
	r = b

	r.ReadBook()

	var w = r.(Write) //因为Book结构体同时实现了Read和Write接口,所以可以这样断言强转类型
	// b: pair<type:Book, value:Book{}地址>
	w.WriteBook()
}
