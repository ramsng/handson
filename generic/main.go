package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type pack interface {
	constraints.Integer | constraints.Float
}

func addfunc[t pack](a t, b t) (c t) {
	c = a + b
	return
}

func main() {
	fmt.Printf("Addition using constraints:%v", addfunc(80.51, 81.49))
}
