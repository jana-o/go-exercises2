//Package acdc asks
package acdc

// Sum adds an unlimited number of ints
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
