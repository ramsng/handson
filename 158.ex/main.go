package main

import "fmt"

//runs in LIFO (LAST IN FIRST OUT) order

func Add(a int, b int) int {
	return a + b
}
func main() {
	a := 5
	b := 10
	fmt.Println("Sum of two integeres", a, b, "is", Add(a, b))
}
