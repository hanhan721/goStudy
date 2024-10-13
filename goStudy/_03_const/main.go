package main

// 常量

const (
	// iota: 关键字,配合const使用,每行的iota都会累加1
	BEIJING = iota
	SHANGHAI
	SHENZHEN
)
const (
	// 可以定义累增规则
	a, b = iota + 1, iota + 3
	c, d
	e, f
)

func main() {
	const length int = 10

}
