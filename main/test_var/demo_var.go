package main

import "fmt"

// 全局变量
var num int = 10

func test() {
	fmt.Printf("num=%d\n", num)
}

// 局部变量
func test1() {
	num := 100
	//var s int = 100
	fmt.Printf("num=%d\n", num) //优先使用局部变量
}

func main() {
	//test()
	test1()
}
