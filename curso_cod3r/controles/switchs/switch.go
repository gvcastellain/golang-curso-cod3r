package main

import "fmt"

func getNotaPorConceito(nota float64) string {
	var notaInt = int(nota)

	switch notaInt { //o switch n precisa de break pq ele por padrao para qnd no case, para contornar isso usar o fllthroug ou o "case ex1, ex2"
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	default:
		return "nota inválida"
	}
}

func main() {
	fmt.Println(getNotaPorConceito(10))
	fmt.Println(getNotaPorConceito(5.2))
}
