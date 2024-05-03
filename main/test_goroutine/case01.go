package main

import (
	"fmt"
	"time"
)

func shopping(name string) {
	fmt.Printf("%s 开始购物 \n", name)
	time.Sleep(1 * time.Second) //定义一秒睡眠时间
	fmt.Printf("%s 结束购物 \n", name)
}

func main() {
	startTime := time.Now()
	shopping("张三")
	shopping("李四")
	shopping("王二")
	shopping("赵五")
	fmt.Println("购买完成 \n", time.Since(startTime))

}
