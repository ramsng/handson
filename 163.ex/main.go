package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//	var x int
	x := func() int {
		x := rand.Intn(10)
		if x == 0 {
			x++
		}
		return x
	}()
	fmt.Println("** Best of luck! Its : ", x)
	fmt.Println("** Lets's see ur second chance! Its : ", y())
}

var y = func() int {
	x := rand.Intn(10)
	if x == 0 {
		x++
	}
	return x
}

func init() {
	fmt.Println("** Let see ur lucky number **")
}
