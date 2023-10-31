package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2

	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notafinal := int(nota)
	fmt.Println(notafinal)

	fmt.Println("convertendo pra string " + strconv.Itoa(123))

	num, _ := strconv.Atoi("123")
	fmt.Println(num)
	fmt.Println(num - 123)

	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("verdadeiro")
	}
}
