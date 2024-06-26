package main

import (
	"fmt"
)

func HRF() bool {
	return true
}

func GFC() int {
	return 3
}




func main() {
	price := 35000

	if price > 50000 {
		if HRF() {
			fmt.Println("신발끈")
		} else{
			fmt.Println("나눠내")
		}
	} else if price >= 30000 && price <= 50000{
		if HRF() {
			fmt.Println("어 신발끈")
		} else{
			fmt.Println("사람 없 나눠내")
		}
	} else {
		fmt.Println("내가 쏨")
	}
}