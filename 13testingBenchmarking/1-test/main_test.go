package main

import "testing"

func TestMySum(t *testing.T) {
	var v int

	v = mySum(1, 2)
	if v != 3 {
		t.Error("expected 3, got", v)
	}
}

//capital MySum
//t.Error equivalent to fmt.Println => no %v

//run command: go test -v
