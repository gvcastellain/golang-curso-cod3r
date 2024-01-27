package main

import "fmt"

func fatorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := fatorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, erro := fatorial(5)
	fmt.Println(resultado, erro)

	_, err := fatorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
