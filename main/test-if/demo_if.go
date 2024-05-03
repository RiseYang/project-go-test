package main

import "fmt"

func main() {

	a := 4
	b := true
	if a > 5 {
		fmt.Println("a>5")
	} else if b == false {
		fmt.Println("b is true")
	} else {
		fmt.Println("final")
	}

}
