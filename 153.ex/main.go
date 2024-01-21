package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "Happy Gophers!"
}

func main() {
	fmt.Println("Foo's Value : ", foo())
	y, x := bar()
	fmt.Println("Bar's Value 1: ", y)
	fmt.Println("Bar's Value 2: ", x)
}
