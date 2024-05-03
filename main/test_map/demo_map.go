package main

import "fmt"

func main() {
	//map声明
	//var man map[string]string
	man := make(map[string]string)
	man["name"] = "zs"
	man["age"] = "18"
	man["sex"] = "男"
	fmt.Printf("man: %v\n", man)

	//第二种赋值方式
	man2 := map[string]string{
		"name": "zs",
		"age":  "18",
		"sex":  "男",
	}
	fmt.Printf("man2: %v\n", man2)

	//ok 判断map的值 true false
	var str = map[string]string{
		"name": "zs",
		"age":  "18",
		"sex":  "男",
	}
	var a1 = "name"
	var b1 = "age1"
	v, ok := str[a1]
	fmt.Printf("v: %v\n", v)
	fmt.Printf("ok: %v\n", ok)
	v1, ok := str[b1]
	fmt.Printf("v: %v\n", v1)
	fmt.Printf("ok: %v\n", ok)

	//map遍历
	var stu = map[string]string{
		"name": "zs",
		"age":  "18",
		"sex":  "男",
	}
	//遍历k的值
	for k := range stu {
		fmt.Printf("k: %v\n", k)
	}
	//遍历k v的值
	for k, v := range stu {
		//fmt.Printf("k: %v\n", k)
		//fmt.Printf("v: %v\n", v)
		fmt.Println("------------")
		fmt.Printf("%v  %v\n", k, v)
	}

}
