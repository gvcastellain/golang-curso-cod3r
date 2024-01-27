package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	verificaPalavra := strconv.Itoa(x)
	var aoContrario string

	for i := len(verificaPalavra) - 1; i >= 0; i-- {
		aoContrario += string(verificaPalavra[i])
	}

	fmt.Println(aoContrario)

	if aoContrario == strconv.Itoa(x) {
		return true
	}

	return false
}

func main() {
	fmt.Println(isPalindrome(121))
}
