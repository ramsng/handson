package main

import "fmt"

func seq1() func() string {
	fmt.Println("Call from Seq 1")
	return seq2
}
func seq2() string {
	return fmt.Sprintln("Call from Seq 2")
}
func seq3() func() string {
	fmt.Println("Call from Seq 3")
	return seq4
}
func seq4() string {
	return fmt.Sprintln("Call from Seq 4")
}

func main() {
	fmt.Println(seq1()())
	fmt.Println(seq3()())
}
