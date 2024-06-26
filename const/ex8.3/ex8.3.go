package main

import (
	"fmt"
)

const (
	// Pig     int = 0
	// Cow     int = 1
	// Chicken int = 2
	// Pig     int = iota //iota = 0부터 1씩 증가
	// Cow     int = iota
	// Chicken int = iota
	Pig     uint = 1 << iota //하나만 선언하면 다음도 규칙 자동 적용, 1 << 0
	Cow                      // 1 << 1
	Chicken                  // 1 << 2
)

func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Pig) //상수명으로 해도
	PrintAnimal(1)   //코드값으로 해도 실행 됨
	PrintAnimal(2)
}
