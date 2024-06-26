package main

import (
	"fmt"
)

type ColorType int 

const ( //구조체
	red ColorType = iota
	blue
	green
	yellow
)

func colorToString(color ColorType) string { //문자열 출력 함수
	switch color {
	case red:
		return "red"
	case blue:
		return "blue"
	case yellow:
		return "yellow"
	default:
		return "undefined"
	}
}

func getMyFavoriteColor() ColorType {
	return blue
}

func main() {
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))
}
