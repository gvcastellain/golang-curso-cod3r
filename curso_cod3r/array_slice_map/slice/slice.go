package main

import (
	"fmt"
	"reflect"
)

func main() {
	arrayPrincipal := [3]int{1, 2, 3}
	slicePrincipal := []int{1, 2, 3}

	fmt.Println(arrayPrincipal, slicePrincipal)
	fmt.Println(reflect.TypeOf(arrayPrincipal), reflect.TypeOf(slicePrincipal))

	arrayExemplo := [5]int{1, 2, 3, 4, 5}
	slicePicotado := arrayExemplo[0:3]
	fmt.Println(arrayExemplo, slicePicotado)
	sliceDiferente := arrayExemplo[:1]
	fmt.Println(sliceDiferente)
}
