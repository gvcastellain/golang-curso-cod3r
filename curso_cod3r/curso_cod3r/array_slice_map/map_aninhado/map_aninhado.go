package main

import "fmt"

func main() {
	funcSalario := map[string]map[string]float64{
		"G": {
			"gabriel":  1500,
			"gabriela": 2500,
		},
		"J": {
			"joão":  9200,
			"joana": 9500,
		},
		"P": {
			"paulo": 53000,
		},
	}

	delete(funcSalario, "P")

	for letra, funcionarios := range funcSalario {
		fmt.Println(letra, funcionarios)
	}

}
