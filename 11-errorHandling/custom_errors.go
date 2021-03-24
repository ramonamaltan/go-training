package main

import (
	"errors"
	"log"
)

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")

func main() {
	// fmt.Printf("%T\n", ErrNorgateMath)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// return 0, ErrNorgateMath
		return 0, errors.New("norgate math: square of negative number")
		// return 0, fmt.Errorf("norgate math: square of negative number %v", f)
	}
	return 42, nil
}
