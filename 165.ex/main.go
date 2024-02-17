package main

import "fmt"

func caller1() func() {
	fmt.Println("Return from caller 1")
	return caller2
}

func caller2() {
	fmt.Println("Return from caller 2")
}

func caller3() func() int {
	fmt.Println("Return from caller 3")
	return func() int {
		fmt.Println("Return from caller 4")
		return 42
	}
}

func main() {
	var y = caller1()
	y()
	var x = caller3()
	x()
	fmt.Println("Final value is:", x())
	fmt.Println("Final value is:", x())
}
