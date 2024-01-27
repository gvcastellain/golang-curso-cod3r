package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "notebook", 1200, []string{"Promoção", "Eletrônicos"}}
	json1, _ := json.Marshal(p1)
	fmt.Println(string(json1))

	var p2 produto
	jsonString := `{"id":2,"nome":"Caneta", "preco":8.90,"tags":["Papelaria","Importado"]}`
	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2.Tags[0])
}
