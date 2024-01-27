package main

import "fmt"

func main() {
	fmt.Print("Mesa")
	fmt.Print(" Linha")

	fmt.Println(" Nova")
	fmt.Println("Linha")

	x := 12.25

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)
	fmt.Println("o valor de x é", x)

	fmt.Printf("o valor de x é %.2f", x)

	a := 1
	b := 1.312
	c := false
	d := "msg"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
