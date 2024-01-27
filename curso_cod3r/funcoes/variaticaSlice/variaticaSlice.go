package main

import "fmt"

func getAprovados(aprovados ...string) {
	fmt.Println("Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d - %s\n", i+1, aprovado)
	}
}

func main() {
	getAprovados("Luiz", "Pedro")
}
