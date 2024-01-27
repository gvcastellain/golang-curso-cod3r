package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("bom dia")
	case t.Hour() > 12:
		fmt.Println("boa tarde")
	}
}
