package main

import "fmt"

func porTipo(tipos interface{}) string {
	switch tipos.(type) {
	case int:
		return "int"
	case float32, float64:
		return "real"
	case string:
		return "string"
	case func():
		return "função"
	default:
		return "desconhecido"
	}
}

func main() {
	fmt.Println(porTipo(33.2))
	fmt.Println(porTipo("opa"))
	fmt.Println(porTipo(1))
	fmt.Println(porTipo(func() {}))
	fmt.Println(porTipo(true))
}
