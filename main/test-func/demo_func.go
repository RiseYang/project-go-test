package main

import (
	"fmt"
	"strings"
)

func test() {
	fmt.Println("hello world")
}

// 定义接收多个参数的函数并且有一个返回值
func test1(a int, b int) int {
	res := a + b
	return res
}

func test2(x, y int) int {
	return x + y
}

// 可变参数
func test3(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数和可变参数同时出现,可变参数要放后面
func test4(a int, b ...int) int {
	sum := a
	for _, v := range b { // _ 表示只读取value的值
		sum = sum + v
	}
	return sum
}

// 定义具有多个返回值
func test5(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sum, sub
}

func main() {
	//函数调用
	res := test1(10, 20)
	fmt.Printf("res: %v\n", res)
	//test()
	s := test3()
	fmt.Printf("s: %v\n", s)
	s1 := test3(10)
	fmt.Printf("s1: %v\n", s1)
	s2 := test3(10, 20)
	fmt.Printf("s2: %v\n", s2)
	s3 := test3(10, 20, 30)
	fmt.Printf("s3: %v\n", s3)

	s4 := test4(1, 2, 3, 4)
	fmt.Printf("s4: %v\n", s4)

	x, y := test5(20, 10)
	fmt.Println(x, y)

	//匿名函数调用
	sayHello := func() {
		fmt.Println("hello world")
	}
	sayHello()

	//第二种调用
	func() {
		fmt.Println("hello world2")
	}()

	//闭包=函数+引用环境(外层变量的引用)
	//调用str函数
	r := str()
	r() //执行str函数内部的匿名函数

	//调用suffix
	strName := makeNameSuffix(".txt")
	allName := strName("hello")
	fmt.Println(allName)
}

// 匿名函数和闭包  闭包=函数+引用环境(外层变量的引用)
// 定义一个函数它的返回值是一个函数
func str() func() {
	name := "zs"
	return func() {
		fmt.Println("hello world3!", name)
	}
}

func makeNameSuffix(suffix string) func(name string) string {
	return func(name string) string {
		//if判断是否存在后缀
		if !strings.HasSuffix(name, suffix) {
			name = name + suffix
		}
		return name
	}
}
