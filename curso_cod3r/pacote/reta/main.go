package main

import "fmt"

func main() {
	p1 := Ponto{2.0, 2.0}
	p2 := Ponto{2.0, 4.0}

	fmt.Println(catetos(p1, p2))
	fmt.Println(distancia(p1, p2))
}

// go run reta.go main.go