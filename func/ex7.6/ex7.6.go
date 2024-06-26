package main

import (
	"fmt"
)

func printNo(n int) { //재귀함수
	if n == 0 { // 탈출 조건
		return
	}

	fmt.Println(n)
	printNo(n - 1)
	fmt.Println("After", n)

}

func main() {
	printNo(3)
}
