package main

import "testing"

func TestSum(t *testing.T) {
	t.Log("sumTest")
	if Sum(1, 2) != 3 {
		t.Error("Sum(1, 2) != 3")
	}
}
