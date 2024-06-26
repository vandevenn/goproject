package main

import (
	"fmt"
)

type House struct {
	Address string
	Size    int
	Price   float64
	Type    string
}

func main() {
	var house House = House{
		"서울시 강동구",
		28,
		9.8,
		"아파트", //마지막 쉼표
	} // 한번에 초기화

	// house.Address = "서울시 강동구"
	// house.Size = 28
	// house.Price = 9.8
	// house.Type = "아파트"

	fmt.Println("주소 :", house.Address)
	fmt.Printf("크기 : %d평\n", house.Size)
	fmt.Printf("가격 : %.2f억원\n", house.Price)
	fmt.Println("타입 :", house.Type)
}


