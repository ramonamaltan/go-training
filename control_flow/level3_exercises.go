package main

import "fmt"

func main() {
	// for statement with ForClause
	// for i := 65; i < 91; i++ {
	// 	fmt.Println(i)
	// 	for j := 1; j < 4; j++ {
	// 		fmt.Printf("\t%#U\n", i)
	// 	}
	// }

	// for condition {}
	year := 1995
	for year < 2021 {
		fmt.Println(year)
		year++
	}

	// for {}
	y := 1995
	for {
		if y > 2020 {
			break
		}
		fmt.Println(y)
		y++
	}

	for i := 10; i < 101; i++ {
		fmt.Printf("When %v divided by 4, the remainder is %v\n", i, i%4)
	}

	y = 1995
	if y == 1990 {
		fmt.Println(y)
	} else if y == 1995 {
		fmt.Println(y)
	} else {
		fmt.Println("not 1990")
	}

	switch {
	case y == 1990:
		fmt.Println(y)
	case y == 1995:
		fmt.Println(y)
	}

	favSport := "spinning"
	switch favSport {
	case "spinning":
		fmt.Println("your fave sport is", favSport)
	case "running":
		fmt.Println("your fave sport is", favSport)
	}
}
