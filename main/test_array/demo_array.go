package main

import "fmt"

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("a: %v\n", a)

	var b [3]int    //定义一个int类型的数组b,长度为3
	var s [2]string //定义一个字符串数组s,长度为2
	//数组初始化
	var b1 = [3]int{1, 2, 3}
	var s1 = [2]string{"hello", "world"}
	var b2 = [2]bool{true, false}
	fmt.Printf("b: %T\n", b) //T表示输出值的类型
	fmt.Printf("s: %T\n", s)
	fmt.Printf("b1: %v\n", b1)
	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("b2: %v\n", b2)

	//数组初始化 ...表示省略长度
	var test = [...]int{1, 2, 3}
	test[0] = 100 //修改长度的值
	fmt.Printf("test: %v\n", test)
	fmt.Printf("test: %v\n", len(test)) //打印长度为3

}
