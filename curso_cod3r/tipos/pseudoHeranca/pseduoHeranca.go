package main

import "fmt"

type cachorro struct {
	raca  string
	cor   string
	idade int
}

type bento struct {
	cachorro
	gostaPassear bool
}

func main() {
	dog := bento{}
	dog.raca = "maltes"
	dog.cor = "branco"
	dog.gostaPassear = true
	dog.idade = 5

	fmt.Println(dog)
}
