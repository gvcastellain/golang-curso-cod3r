package main

import (
	"fmt"
	"math"
)

func main() {
	a := 2
	b := 3

	fmt.Println("soma: ", a+b)
	fmt.Println("sub: ", a-b)
	fmt.Println("mutli: ", a*b)
	fmt.Println("divi: ", a/b)
	fmt.Println("resto: ", a%b)

	fmt.Println("e: ", a&b)
	fmt.Println("ou: ", a|b)
	fmt.Println("xor: ", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("maior", math.Max(float64(a), float64(b)))
	fmt.Println("maior", math.Max(c, d))
	fmt.Println("espo", math.Pow(c, d))
}
