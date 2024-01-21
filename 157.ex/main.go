package main

import (
	"fmt"
	"math"
)

// Hands-on exercise #62 - interfaces
// ****************************************************************************************************************************
// Keypoint-takeaway - same function name can be used for 2 different methods
// Keypoint-takeaway - Return should be defined in an interface type; function associated with interface should also be returned
// ****************************************************************************************************************************
// ● create a type SQUARE
// ○ length float64
// ○ width float64
// ● create a type CIRCLE
// ○ radius float64
// ● attach a method to each that calculates AREA and returns it
// ○ circle area= π r 2
// ■ math.Pi
// ■ math.Pow
// ○ square area = L * W
// ● create a type SHAPE that defines an interface as anything that has the AREA method
// ● create a func INFO which takes type shape and then prints the area
// ● create a value of type square
// ● create a value of type circle
// ● use func info to print the area of square
// ● use func info to print the area of circle
type square struct {
	length float64
	width  float64
}
type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
	//	areacircle()
}

func info(s shape) float64 {
	return s.area()
}

/*
	func info(shape string, len int, width int) {
		switch shape {
		case "square":
			fmt.Println("Shape: ", shape, "\n Area: ", areasqr())
		case "circle":
			fmt.Println("Shape: ", shape, "\n Area: ", areacircle())
		default:
			fmt.Println("Shape: ", shape, "\n Area not available ")
		}
	}
*/
func main() {
	q1 := square{
		length: 4,
		width:  4,
	}
	q2 := circle{
		radius: 4,
	}
	fmt.Println("Info q1", info(q1))
	fmt.Println("Info q1", info(q2))
	//fmt.Println("Info q2", q1.area())
	//fmt.Println("Info q2", q2.area())
	//areaxface(q1)
	/*	q2 := circle{
		b := areaxface(q1)
			radius: 4,
		}
		fmt.Println("Info q2", q2.areacircle(q2.radius))
		a := areaxface(q1)
		fmt.Println("Print through interfaces", a)
	*/
}
