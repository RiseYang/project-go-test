package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	//字符串连接
	var name string = "张三"
	var school string = "石牛小学"

	msg := name + school
	fmt.Printf("msg: %v\n", msg)

	//buffer做字符串连接
	var buffer bytes.Buffer
	buffer.WriteString("hello")
	buffer.WriteString(" ")
	buffer.WriteString("world")
	fmt.Printf("buffer.String(): %v\n", buffer.String())

	//字符串函数
	var str string = "hello World"
	println(len(str))                                                      //字符串长度
	fmt.Printf("strings.Split(str, \" \"): %v\n", strings.Split(str, " ")) //字符串分割
	println(strings.Split(str, " "))
	fmt.Printf("strings.toupper: %v\n", strings.ToUpper("hello")) //字符串大写
	fmt.Printf("strings.tolower: %v\n", strings.ToLower("world")) //字符串小写
	print(strings.Contains(str, "hello"))                         //判断字符串是否包含
	print(strings.HasPrefix(str, "hello"))
	println(strings.HasSuffix(str, "world"))
	println(strings.Index(str, "d"))
	//字符串转义符
	var human string = "hello world"
	fmt.Printf("human:\n %v\n", human) //换行符
	fmt.Printf("human:\t %v\n", human) //制表符
	fmt.Printf("human:\r %v\n", human) //回车符

	//字符串切片操作
	s := "hello world"
	a := 1
	b := 5
	fmt.Printf("a: %c\n", s[a])     //%c表示一个字符
	fmt.Printf("a:b: %v\n", s[a:b]) //从a到b
	fmt.Printf("a: %v\n", s[a:])    //从a到最后
	fmt.Printf(":b %v\n", s[:b])    //开头到b

}
