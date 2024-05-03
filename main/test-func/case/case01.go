package _case

import "errors"

//形参  (a, b int)
//实参

func SumCase(a, b int) (num int, err error) {

	if a <= 0 && b <= 0 {
		err = errors.New("两数相加不能小于0")
	}
	num = a + b
	return
}
