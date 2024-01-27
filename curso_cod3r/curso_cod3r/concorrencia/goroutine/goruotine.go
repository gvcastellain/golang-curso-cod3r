package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("joao", "oi", 3)
	// fale("paulo", "retorno do oi", 1)

	go fale("ricardo", "eai", 10)
	go fale("cata", "opa", 10)

	go fale("pessoa", "texto", 5)
	fale("joao", "frase", 3)

	//qnd coloca o go na frente o processo vai pra uma thread "secundaria"
	//enqt o resto do programa ainda esta na thread principal e acaba apos ser executada, sem terminar a func anterior
	//usar os chanel resolve isso
}
