package main

import (
	"fmt"
	"math"
)

func overflow() {
	var x int8 = 127

	fmt.Printf("%d < %d + 1: %v\n", x, x, x < x+1)
	fmt.Printf("x\t= %4d, %08b\n", x, x)
	fmt.Printf("x + 1\t= %4d, %08b\n", x+1, x+1)
	fmt.Printf("x + 2\t= %4d, %08b\n", x+2, x+2)
	fmt.Printf("x + 3\t= %4d, %08b\n", x+3, x+3)

	var y int8 = -128

	fmt.Printf("%d > %d - 1: %v\n", y, y, y > y-1)
	fmt.Printf("y\t= %4d, %08b\n", y, y)
	fmt.Printf("y - 1\t= %4d, %08b\n", y-1, y-1)
}

func float() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f : %v\n", a, b, c, a+b == c)
	fmt.Println(a + b)
}

const epsilon = 0.000001

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func ignore() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3


	
	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}

func main() {

	//overflow()
	//float()
	ignore()
}
