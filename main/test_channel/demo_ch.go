package main

import "fmt"

// 声明一个有缓冲的 chan(管道)
func main() {
	fmt.Println("main start")
	ch := make(chan string, 1)
	ch <- "a" //入chan
	go func() {
		val := <-ch //出chan
		fmt.Println(val)
	}()
	fmt.Println("main end")
}
