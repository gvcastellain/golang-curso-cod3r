package main

import "fmt"

func multiplica(a, b int) int {
	return a * b
}

func exec(funcao func(int, int) int, p1, p2 int) int { //func q recebe outra func q tem como assinatura 2 params int e retorno int
	return funcao(p1, p2)
}

func main() {
	resultado := exec(multiplica, 3, 4)
	fmt.Println(resultado)
}
