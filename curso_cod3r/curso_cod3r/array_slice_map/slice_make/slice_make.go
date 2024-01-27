package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s[9])

	sliceCapacidade := make([]int, 10, 20)
	fmt.Println(sliceCapacidade, len(sliceCapacidade), cap(sliceCapacidade))
	sliceCapacidade = append(sliceCapacidade, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(sliceCapacidade, len(sliceCapacidade), cap(sliceCapacidade))

	sliceCapacidade = append(sliceCapacidade, 312)
	fmt.Println(sliceCapacidade, len(sliceCapacidade), cap(sliceCapacidade))

}
