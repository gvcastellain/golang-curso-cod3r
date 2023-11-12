package main

import "fmt"

func printaAlgo(palavraPrintada string) {
	fmt.Println(palavraPrintada)
}

func retornaStrings() (string, string) {
	return "palavra 1", "palavra 2"
}

func main() {
	printaAlgo("opa")

	retorno1, _ := retornaStrings()
	fmt.Println(retorno1)
	retorno1, retorno2 := retornaStrings()
	fmt.Println(retorno1, retorno2)
}
