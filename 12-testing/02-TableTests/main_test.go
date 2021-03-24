package main

import (
	"testing"
)

func TestMySum(t *testing.T) {
	// table tests: writing a series of tests to run. Allows a variety of situations
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{30, 12}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 1}, 2},
		test{[]int{-1, 0, 1}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}
}
