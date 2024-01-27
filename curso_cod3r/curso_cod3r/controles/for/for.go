package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1

	for i <= 10 {
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println()

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("par: %d ", i)
		} else {
			fmt.Printf("ímpar: %d ", i)
		}
	}

	for { //laco infinito
		fmt.Println("for infinito")
		time.Sleep(time.Second)
	}
}
