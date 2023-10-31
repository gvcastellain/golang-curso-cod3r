package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.gsuplementos.com.br/whey-protein-concentrado-1kg-growth-supplements-p985936")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
