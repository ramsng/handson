package main

import "testing"

func TestParadise(t *testing.T) {
	x := Paradise("Enjoy the food")
	y := "Paradise Quote:Enjoy the food"
	if x != y {
		t.Errorf("error \n%v\n%v", x, y)
	}
}
