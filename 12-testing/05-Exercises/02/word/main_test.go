package word

import (
	"fmt"
	"testing"

	"github.com/training/12-testing/05-Exercises/02/quote"
)

func TestCount(t *testing.T) {
	s := "This is a string"
	c := Count(s)
	if c != 4 {
		t.Error("got ", c, "wanted 4")
	}
}

func TestUseCount(t *testing.T) {
	s := "One two three three three"
	m := UseCount(s)
	for k, v := range m {
		switch k {
		case "One":
			if v != 1 {
				t.Error("got ", v, "wanted 1")
			}
		case "two":
			if v != 1 {
				t.Error("got ", v, "wanted 1")
			}
		case "thre":
			if v != 3 {
				t.Error("got ", v, "wanted 3")
			}
		}
	}
}

func ExampleCount() {
	s := "This is a string"
	fmt.Println(Count(s))
	// Output:
	// 4
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count(quote.SunAlso)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(quote.SunAlso)
	}
}
