package main

import (
	"fmt"
)

func getMyAge() (int, bool) {
	return 33, true
}

func main() {

	var age2 int

	if age, ok := getMyAge(); ok && age < 20 {
		fmt.Println("you are young", age)
		age2 = age
	} else if ok && age < 30 {
		fmt.Println("Nice age", age)
		age2 = age
	} else if ok {
		fmt.Println("you are beautiful", age)
		age2 = age
	} else {
		fmt.Println("Error")
		age2 = age
	}

	fmt.Println("your age is", age2)
}
