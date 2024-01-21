package main

import "fmt"

//runs in LIFO (LAST IN FIRST OUT) order

func foo() {
	fmt.Println("EXECUTE 4 => Foo execute")
}
func bar() {
	defer fmt.Println("EXECUTE 3 => Bar1 execute")
	fmt.Println("EXECUTE 2 => Bar2 execute")
	defer foo()
}
func main() {
	fmt.Println("EXECUTE 1 => 155. DEFER FUNCTION EXERCISE")
	defer fmt.Println("EXECUTE 5 =>  DEFER PRINT 1")
	defer foo()
	defer bar()
}
