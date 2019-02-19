package word

import (
	"exercises2/13testingBenchmarking/exercises/2-exercises/quote"
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	m := UseCount("one two three three three")
	for k, v := range m {
		switch k {
		case "one":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "two":
			if v != 1 {
				t.Error("got", v, "want", 1)
			}
		case "three":
			if v != 3 {
				t.Error("got", v, "want", 3)
			}
		}
	}
}

func TestCount(t *testing.T) {
	n := Count("Hello World")
	if n != 2 {
		t.Error("want", 2, "got", n)
	}
}

func ExampleCount() {
	fmt.Println(Count("Hello World"))
	//Output: 2
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}

}
