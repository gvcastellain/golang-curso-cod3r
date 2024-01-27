package main

import "fmt"

func main() {
	funcSalarios := map[string]float64{
		"jose":    1500,
		"andreia": 4600.50,
		"Pedro":   3200,
	}

	funcSalarios["rafaela"] = 7250
	delete(funcSalarios, "n tem")

	for nome, salario := range funcSalarios {
		fmt.Println(nome, salario)
	}
}
