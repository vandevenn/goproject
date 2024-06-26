package main

import (
	"fmt"
)

func main() {
	var choice int
	var num1, num2 float64

	for {
		fmt.Println("계산기")
		fmt.Println("1. 덧셈")
		fmt.Println("2. 뺄셈")
		fmt.Println("3. 곱셈")
		fmt.Println("4. 나눗셈")
		fmt.Println("5. 종료")
		fmt.Print("계산 선택: ")
		fmt.Scanln(&choice)

		if choice == 5 {
			fmt.Println("프로그램을 종료합니다")
			break
		}

		fmt.Print("첫 번째 숫자를 입력하세요: ")
		fmt.Scanln(&num1)
		fmt.Print("두 번째 숫자를 입력하세요: ")
		fmt.Scanln(&num2)

		switch choice {
		case 1:
			fmt.Printf("결과: %.2f\n", add(num1, num2))
		case 2:
			fmt.Printf("결과: %.2f\n", sub(num1, num2))
		case 3:
			fmt.Printf("결과: %.2f\n", multiply(num1, num2))
		case 4:
			if num2 != 0 {
				fmt.Printf("결과: %.2f\n", divide(num1, num2))
			} else {
				fmt.Printf("0으로 나눌 수 없습니다.")
			}
		default:
			fmt.Println("잘못된 선택입니다.")

		}




	}
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {


	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
