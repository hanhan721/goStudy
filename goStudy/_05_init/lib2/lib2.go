package lib2

import "fmt"

func init() {
	fmt.Println("lib2.init().....")
}

// Test 方法名首字母大写才能被其他包调用
func Test() {
	fmt.Println("lib2 test.....")
}
