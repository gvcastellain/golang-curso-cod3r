package main

import "fmt"

func main() {
	i := 1
	l := 2

	var p *int = nil

	p = &i //pega o endereco de memoria que aponta pro i

	fmt.Println(p, *p, i, &i)
	*p++ //nao da pra dar p++
	i++
	fmt.Println(p, *p, i, &i)

	p = &l
	fmt.Println(p)
}
