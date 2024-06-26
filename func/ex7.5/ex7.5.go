package main

import "fmt"

func Divide(a, b int) (result int, success bool) { // 반환할 변수명 명시, 무엇을 return 할지 지정안해도 됨
	if b == 0 {
		result = 0
		success = false
		return
	} else {
		result = a / b
		success = true
		return
	}
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
