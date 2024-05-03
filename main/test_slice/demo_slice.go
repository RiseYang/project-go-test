package main

import "fmt"

func main() {
	var names []string
	var numbers []int

	fmt.Printf("names: %v\n", names)
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Println(names == nil)
	fmt.Println(numbers == nil)

	//make声明
	var str []string = make([]string, 0)
	var str2 = make([]int, 2) //声明2个长度的切片
	fmt.Printf("str: %v\n", str)
	fmt.Printf("str2: %v\n", str2)

	//切片长度len和容量cap
	var test1 = []string{"zs", "ls"}
	var test2 = []int{1, 2, 3}
	//打印值
	fmt.Printf("test1: %v\n", test1)
	fmt.Printf("test2: %v\n", test2)
	//打印长度和容量
	fmt.Printf("len: %d cap: %d\n", len(test1), cap(test1))
	fmt.Printf("len: %d cap: %d\n", len(test2), cap(test2))
	//打印切片的长度和容量
	var s1 = make([]string, 2, 3)
	fmt.Printf("len: %d cap: %d\n", len(s1), cap(s1))

	//切片初始化
	m := []int{1, 2, 3}
	n := []string{"hello", "world"}
	fmt.Printf("m: %v\n", m)
	fmt.Printf("n: %v\n", n)

	//第二种数组初始化
	arr := [...]int{1, 2, 3}
	d := arr[:] //arr[:]表示所有数组 省略的写法
	fmt.Printf("d: %v\n", d)

	arr2 := [...]int{1, 2, 3, 4, 5, 6}
	d1 := arr2[2:5]            //[2:5]=>[)区间
	fmt.Printf("d1: %v\n", d1) //切片索引范围左包含,右不包含
	d2 := arr2[2:]
	fmt.Printf("d2: %v\n", d2)
	d3 := arr2[:3]
	fmt.Printf("d3: %v\n", d3)

	//切片遍历
	var num = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < len(num); i++ {
		fmt.Printf("f[%d]: %v\n", i, num[i]) //%d 表示十进制整数
	}

	//for range遍历
	stu := []string{"hello", "world"}
	for i, v := range stu {
		fmt.Printf("i: %v\n", i) //i表示索引
		fmt.Printf("v: %v\n", v) //v表示值
		fmt.Printf("i: %d %v\n", i, v)
	}
}
