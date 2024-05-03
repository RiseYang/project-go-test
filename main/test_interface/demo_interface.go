package main

import "fmt"

type Dog struct {
}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

// 定义一个抽象类型,只要实现say()方法类型都是sayer类型
type sayer interface {
	say()
}

func Call(arg sayer) {
	arg.say()
}

type Cat struct {
}

func (c Cat) say() string {
	return "喵喵喵..."
}

func main() {
	//调用call
	d := Dog{}
	Call(d)

	c := Cat{}
	fmt.Println("cat:", c.say())
}
