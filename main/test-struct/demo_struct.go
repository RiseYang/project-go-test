package main

import "fmt"

type student struct {
	name string
	age  int
	sex  string
}

func main() {
	//声明变量
	var str student
	str = student{name: "张三", age: 18, sex: "男"}
	fmt.Println(str)
	fmt.Printf("str: %v\n", str)
	fmt.Printf("str: %+v\n", str) //打印详情
	fmt.Printf("str: %#v\n", str) //打印main函数下的详细信息
}
