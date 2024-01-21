package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) speak() {
	fmt.Println("Name :", p.name, "\nAge  :", p.age)
}

func main() {
	p1 := person{
		name: "Alan",
		age:  42,
	}
	p2 := person{
		name: "Donald",
		age:  24,
	}
	p1.speak()
	p2.speak()
}
