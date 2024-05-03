package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//创建打开当前目录下的main.go文件
	file, err := os.Open("./main/test_file/test.txt")
	if err != nil {
		fmt.Printf("打开文件失败!, err:%v\n", err)
		return
	}
	//关闭文件
	defer file.Close()

	for {
		//读取文件
		var tmp = make([]byte, 128)
		n, err := file.Read(tmp)
		if err == io.EOF {
			//把读出的当前的文件打印出来
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("读取文件失败! err: %v\n", err)
			return
		}
		fmt.Printf("read success  %d bytes from file. \n", n)
		fmt.Println(string(tmp[:n]))
	}
}
