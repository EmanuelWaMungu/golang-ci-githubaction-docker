package main

import "testing"

func TestMultiply(t *testing.T) {

	if Multiply(2, 3) != 6 {
		t.Error("Expected 6")
	}
}
