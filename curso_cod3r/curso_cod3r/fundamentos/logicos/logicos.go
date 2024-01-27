package main

import "fmt"

func compras(sobrouSetembro, sobrouOutubro bool) (bool, bool, bool) {
	comprarTecladoSemFio := sobrouSetembro && sobrouOutubro
	comprarTecladoPadrao := sobrouSetembro != sobrouOutubro
	comprarFone := sobrouSetembro || sobrouOutubro

	return comprarTecladoSemFio, comprarTecladoPadrao, comprarFone
}

func main() {
	tecladoSemFio, tecladoPadrao, fone := compras(true, true)
	fmt.Printf("teclado sem fio: %t, teclado padrão %t, fone: %t, negação do fone -> %t", tecladoSemFio, tecladoPadrao, fone, !fone)

}
