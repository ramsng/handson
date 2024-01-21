package main

import "testing"

//runs in LIFO (LAST IN FIRST OUT) order

func TestAdd(t *testing.T) {
	total := Add(5, 10)
	if total != 15 {
		t.Errorf("total received :", total, 15)
	}
}
