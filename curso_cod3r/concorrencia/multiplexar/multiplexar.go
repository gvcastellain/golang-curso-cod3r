package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func titulo(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}
	return c
}

func encaminha(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminha(entrada1, c)
	go encaminha(entrada2, c)
	return c
}

func main() {
	c := juntar(
		titulo("https://www.google.com", "https://www.furb.br"),
		titulo("https://www.youtube.com", "https://www.twitch.com"),
	)

	fmt.Println(<-c, " - ", <-c)
	fmt.Println(<-c, " - ", <-c)
}
