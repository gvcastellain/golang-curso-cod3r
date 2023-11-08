package main

import (
	"fmt"
)

func main() { //array de numeros fixos mas conta a partir da incializacao
	numeros := [...]int{1, 2, 3, 4, 5}

	for i, valorNumeros := range numeros { //retorna o indice e o valor
		fmt.Printf("%d) %d\n", i, valorNumeros)
	}

	for _, valorSemIndice := range numeros { //ignora o indice
		fmt.Printf("%d\n", valorSemIndice)
	}
}
