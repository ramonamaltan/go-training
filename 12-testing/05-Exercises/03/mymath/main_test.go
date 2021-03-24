package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	type test struct {
		data   []int
		answer float64
	}

	tests := []test{
		test{[]int{3, 5, 89, 200, 5, 29}, 32},
		test{[]int{45, 6, 90}, 45},
		test{[]int{-50, 100, 3000, -100}, 25},
	}

	for _, v := range tests {
		f := CenteredAvg(v.data)
		if f != v.answer {
			t.Error("got", f, "expected", v.answer)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{45, 6, 90}))
	// Output:
	// 45
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{45, 6, 90})
	}
}
