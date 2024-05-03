package main

import "fmt"

func main() {
	var ptr *int                 //定义一个int类型的指针
	fmt.Printf("ptr: %x\n", ptr) //打印为0 %x表示小写的十六进制数值
	fmt.Printf("ptr: %v\n", ptr) //打印为<nil>
	//空指针判断
	if ptr != nil {
		fmt.Printf("ptr: %x\n", ptr)
	} //ptr 不是空指针
	if ptr == nil {
		fmt.Printf("ptr: %x\n", ptr)
	} // ptr 是空指针

}
