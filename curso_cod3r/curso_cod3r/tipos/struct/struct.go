package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

func (produtoFicticio produto) precoComDesconto() float64 {
	return produtoFicticio.preco * (1 - produtoFicticio.desconto)
}

func main() {
	var produto1 produto
	produto1 = produto{
		nome:     "caixa",
		preco:    10.50,
		desconto: 0.05,
	}

	fmt.Println(produto1)

	produto2 := produto{
		nome:     "mesa",
		preco:    280.75,
		desconto: 0.10,
	}

	fmt.Println(produto2, produto2.precoComDesconto())

	produto3 := produto{"cadeira", 400, 0.2}
	fmt.Println(produto3)
}
