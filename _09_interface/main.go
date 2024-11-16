package main

import "fmt"

// interface实现多态
func main() {
	var animal Animal
	fmt.Printf("type of animal: %T\n", animal)
	animal = &Cat{"黄色的"}
	animal.Sleep()
	fmt.Printf("type of animal: %T\n", animal)
	animal = &Dog{"黑色的"}
	animal.Sleep()
	fmt.Printf("type of animal: %T\n", animal)
	fmt.Println("-------------------")
	ShowAnimal(animal)
	cat := &Cat{"白色的"}
	ShowAnimal(cat)
}
