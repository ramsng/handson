package main

import "fmt"

func varsum(ii ...int) int {
	var a int = 0
	for _, b := range ii {
		a += b
	}
	return a
}

func sum(ii []int) int {
	var a int = 0
	for _, b := range ii {
		a += b
	}
	return a
}

func main() {
	i := []int{1, 2, 3, 4, 54, 6, 7, 8, 8}
	fmt.Println("Sum of Variadic func parms        : ", varsum(i...))
	fmt.Println("Sum of func parms passed in slice : ", sum(i))
}
