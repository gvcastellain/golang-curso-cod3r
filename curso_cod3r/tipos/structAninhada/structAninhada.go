package main

import "fmt"

type item struct {
	produtoID int
	qtd       int
	preco     float64
}

type pedido struct {
	usuarioID int
	itens     []item
}

func (p pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtd)
	}
	return total
}

func main() {
	pedido := pedido{
		usuarioID: 1,
		itens: []item{
			item{produtoID: 1, qtd: 2, preco: 20},
			item{2, 1, 17.50},
		},
	}

	fmt.Println(pedido.valorTotal())
}
