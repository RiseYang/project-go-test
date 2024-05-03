package main

import "fmt"

func main() {
	fmt.Println("start")
	//defer 定义的逆序进行执行,defer语句会将其后面的进行延迟处理
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	fmt.Println("end")
}
