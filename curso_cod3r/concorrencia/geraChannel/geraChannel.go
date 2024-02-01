package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func main() {
	t1 := titulo("https://www.google.com", "https://www.furb.br")
	t2 := titulo("https://www.youtube.com", "https://www.twitch.com")

	time.Sleep(time.Second * 4)

	fmt.Println(<-t1, " -- ", <-t2)
	fmt.Println(<-t1, " -- ", <-t2)
}
