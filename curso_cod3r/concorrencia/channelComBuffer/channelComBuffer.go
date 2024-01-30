package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	fmt.Println("RODOU")
	ch <- 3
	fmt.Println("RODOU")
	ch <- 4
	fmt.Println("NAO VAI")
	ch <- 5
}

func main() {
	ch := make(chan int, 2)
	go rotina(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
