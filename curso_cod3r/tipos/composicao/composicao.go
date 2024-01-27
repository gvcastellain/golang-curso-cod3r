package main

import "fmt"

type tecladoMecanico interface {
	ligaLed()
}

type tecladoBluetooth interface {
	ligaBluetooth()
}

type telcadoMecanicoBluetooh interface {
	tecladoMecanico
	tecladoBluetooth
}

type tecladoRedragon struct{}

func (t tecladoRedragon) ligaLed() {
	fmt.Println("ligando os leds")
}

func (t tecladoRedragon) ligaBluetooth() {
	fmt.Println("ligando o bluetooth")
}

func main() {
	var teclado telcadoMecanicoBluetooh = tecladoRedragon{}
	teclado.ligaBluetooth()
	teclado.ligaLed()
}
