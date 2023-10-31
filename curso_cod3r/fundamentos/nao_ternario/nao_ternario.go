package main

import "fmt"

func getResultado(nota float64) string {
	if nota >= 6 {
		return "aprovado"
	}

	return "reprovado"
}

func main() {
	fmt.Println(getResultado(6.2))
}
