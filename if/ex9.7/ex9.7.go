package main

import (
	"fmt"
)

func main() {
	var temp int = 30
	var rain int = 40

	if temp >= 25 {
		if rain >= 80 {
			fmt.Println("덥고 비가옵니다.")
		} else if rain >= 20 {
			fmt.Println("덥고 습합니다")
		} else {
			fmt.Println("야외 활동하기 좋습니다")
		}
	} else {
		if temp < 10 || rain >= 80 {
			fmt.Println("야외 활동하기 좋지 않습니다")
		}

		fmt.Println("좋은 날씨 입니다.")
	}
}