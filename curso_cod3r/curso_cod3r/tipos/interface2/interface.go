package main

import "fmt"

type carro interface {
	ligaTurbo()
}

type monza struct {
	modelo          string
	turboLigado     bool
	velocidadeAtual int
}

func (m *monza) ligaTurbo() {
	m.turboLigado = true
}

func main() {
	carro1 := monza{"velho", false, 0}
	carro1.ligaTurbo()

	var carro2 carro = &monza{"velho", false, 0}
	carro2.ligaTurbo()

	fmt.Println(carro1, carro2)
}
