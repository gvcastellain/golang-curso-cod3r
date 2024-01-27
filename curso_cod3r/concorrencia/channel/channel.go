package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 //envia a um canal um dado, nesse caso o canal eh de int, escrita

	<-ch //recebe dados do canal, leitura

	ch <- 2
	fmt.Println(<-ch)
}
