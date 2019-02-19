package mystr

import "strings"

func Cat(xs []string) string {
	s := xs[0]
	for _, v := range xs[1:] {
		s += " "
		s += v
	}
	return s
}

//very resource intensive, slow

func Join(xs []string) string {
	return strings.Join(xs, " ")
}

//run command: go test -bench .

//len of string is number of bytes

//remember to BET (Benchmark, Exampe, Test)
