package main

import "fmt"

func main() {
	//for循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for range循环
	var a = [5]int{1, 2, 3, 4, 5} //声明数组为5
	for i, v := range a {
		//fmt.Printf("i: %d %v\n", i, v) //i为索引 v表示值
		fmt.Printf("i: %v\n", i)
		fmt.Printf("v: %v\n", v)
	}
}
