package main

import (
	"fmt"
	"math"
)

func squareseq() func() float64 {
	var i float64 = 1
	return func() float64 {
		y := math.Pow(2, i)
		i++
		return y
	}
}
func main() {
	fmt.Printf("%T\t%v\n", squareseq()(), squareseq()())
	a := squareseq()
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", squareseq()(), squareseq()())
	a = squareseq()
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())
	fmt.Printf("%T\t%v\n", a(), a())

}

/*Exercise:72 - Closure exercise
package main

import (
	"fmt"
	"math"
	"math/rand"
)

func randvalgen() func() float64 {
	var x float64 = float64(rand.Intn(10))
	fmt.Println("Randomvalue", x+1)

	return func() float64 {
		x++
		return math.Pow(x, 2)
		//	return y
	}
	//	return y
}
func main() {
	y := randvalgen()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	y = randvalgen()
	fmt.Println(y())
	fmt.Println(y())
}

*/
/*Call back */
