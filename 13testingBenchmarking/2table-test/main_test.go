package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{1, 2}, 3},
		test{[]int{21, 21}, 42},
		test{[]int{1, 0, -1}, 0},
	}
	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("expected", v.answer, "got: ", x)
		}
	}
}

// mySum(xi ...int) int  ==> x := mySum(v.data...)
