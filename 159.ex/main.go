package main

//
import "fmt"

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func doMath(a int, b int, c func(int, int) int) int {
	return c(a, b)
}
func main() {
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)

	a := 10
	b := 10
	fmt.Println("Add of ", a, b, ":", add(a, b))
	fmt.Println("Diff of ", a, b, ":", sub(a, b))
	fmt.Println("Math comp of ", a, b, ":", doMath(a, b, add))
	fmt.Println("Math comp of ", a, b, ":", doMath(a, b, sub))
}
