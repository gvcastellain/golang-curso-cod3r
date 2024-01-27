package main

import "fmt"

func trocar(p1, p2 int) (primeiro int, segundo int) {
	segundo = p2
	primeiro = p1
	return
}

func main() {
	x, y := trocar(1, 2)

	fmt.Println(x, y)
}
