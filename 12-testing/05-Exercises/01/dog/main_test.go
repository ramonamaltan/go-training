package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y1 := Years(5)
	if y1 != 35 {
		t.Error("got", y1, "wanted 35")
	}
}

func TestYearsTwo(t *testing.T) {
	y2 := YearsTwo(10)
	if y2 != 70 {
		t.Error("got", y2, "wanted 70")
	}
}

func ExampleYears() {
	fmt.Println(Years(5))
	// Output:
	// 35
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output: 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
