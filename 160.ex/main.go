package main

import "fmt"

func Paradise(s string) string {
	return fmt.Sprintf("Paradise Quote:%v", s)
}
func main() {
	fmt.Println(Paradise("Enjoy the food"))
}
