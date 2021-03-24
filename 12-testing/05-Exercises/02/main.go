package main

import (
	"fmt"

	"github.com/training/12-testing/05-Exercises/02/quote"
	"github.com/training/12-testing/05-Exercises/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
