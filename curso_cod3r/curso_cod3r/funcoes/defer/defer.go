package main

import "fmt"

func getValorAprovado(numero int) int {
	defer fmt.Println("fim")
	if numero > 5000 {
		fmt.Println("valor alto")
		return 5000
	} else {
		fmt.Println("valor baixo")
		return numero
	}
}

func main() {
	getValorAprovado(6000)
	getValorAprovado(3000)
}
