package main

import "testing"

func TestSub(t *testing.T) {
	total := sub(5, 5)
	if total != 10 {
		t.Errorf("Bug in total: %v", total)
	}
}
func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Bug in total: %v", total)
	}

}
