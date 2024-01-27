package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 100)
	fmt.Println("literal inteiro é", reflect.TypeOf(32000)) //pega o parametro e passa o tipo d

	//uint - 8 16 32 64,
	var b byte = 255                           //byte é o apelido do unit8
	fmt.Println("o byte é", reflect.TypeOf(b)) //vem como uint8

	maxI := math.MaxInt64
	fmt.Println("o valor máximo do int é", maxI)
	fmt.Println("o byte é", reflect.TypeOf(maxI))

	var i2 rune = 'a'
	fmt.Println("o rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	var x float32 = 32.23 //var x = float32(32.32)
	fmt.Println("tipo de x", reflect.TypeOf(x))
	fmt.Println("o tipo da var é", reflect.TypeOf(32.32)) //por padrao 64

	bo := true
	fmt.Println("o tipo do bo", reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "passando a string"
	fmt.Println(s1 + "!")
	fmt.Println("tamanho da string", len(s1))

	s2 := `passando
	 a 
	 string`
	fmt.Println(s2 + "!")
	fmt.Println("tamanho da string", len(s2))

}
