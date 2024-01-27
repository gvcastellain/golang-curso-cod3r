package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func handleError(er error) {
	if er != nil {
		panic(er)
	}
}

func linkVisited(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		fmt.Println(node.Data)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		linkVisited(c)
	}
}

func main() {
	url := "https://www.furb.br/dion/aluno/diario.xhtml?faces-redirect=true"

	resp, err := http.Get(url)

	handleError(err)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status diferente de 200 %d", resp.StatusCode))
	}

	doc, err := html.Parse(resp.Body)

	handleError(err)

	linkVisited(doc)
}
