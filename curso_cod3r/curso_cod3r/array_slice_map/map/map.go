package main

import "fmt"

func main() {
	aprovados := make(map[int]string) //chave valor

	aprovados[112312] = "claudio"
	aprovados[141331] = "paula"
	aprovados[213123] = "bento"
	fmt.Println(aprovados)

	for id, nome := range aprovados {
		fmt.Printf("%s (ID: %d) \n", nome, id)
	}

	fmt.Println(aprovados[112312])
	delete(aprovados, 112312)
	fmt.Println(aprovados[112312])
}
