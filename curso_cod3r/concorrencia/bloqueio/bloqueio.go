package main

import (
	"fmt"
	"time"
)

func rotina(c chan int) {
	time.Sleep(time.Second)
	c <- 1
	fmt.Println("só depois que foi lido")
}

func main() {
	c := make(chan int)

	go rotina(c)

	fmt.Println(<-c)
	fmt.Println("Foi lido")
	fmt.Println(<-c)
	fmt.Println("fim")
}
