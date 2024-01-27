package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numeroAleatorio() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)
}

func main() {
	if i := numeroAleatorio(); i > 5 { //inicializa a var no if e faz a operação, cria ela local, da pra fazer no switch tbm
		fmt.Println("i maior que 5, valor:", i)
	} else {
		fmt.Println("i menor que 5, valor:", i)
	}
}
