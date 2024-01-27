package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 8 && nota < 9 {
		return "B"
	} else if nota >= 6 && nota < 8 {
		return "C"
	} else if nota >= 4 && nota < 6 {
		return "D"
	} else if nota >= 2 && nota < 4 {
		return "E"
	} else {
		return "F"
	}
}

func main() {
	fmt.Printf("notas: %v, %v, %v, %v, %v, %v",
		notaParaConceito(10), notaParaConceito(8.2), notaParaConceito(6.1), notaParaConceito(4), notaParaConceito(3.5), notaParaConceito(1))
}
