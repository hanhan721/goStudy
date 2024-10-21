package main

import "fmt"

type Animal interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的具体种类
}

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println(c.GetColor(), c.GetType(), "正在睡觉....")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "小猫"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println(d.GetColor(), d.GetType(), "正在睡觉....")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (c *Dog) GetType() string {
	return "小狗"
}

func ShowAnimal(animal Animal) {
	animal.Sleep()
}
