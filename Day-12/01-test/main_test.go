package main

import "testing"

func TestSum(t *testing.T) {
	x := sumf(25, 7)
	if x != 32 {
		t.Error("Expected", 32, "Got", x)
	}
}
