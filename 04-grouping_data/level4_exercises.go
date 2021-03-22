package main

import "fmt"

func main() {
	// array
	x := [5]int{}
	fmt.Println(x)

	x[0] = 8
	x[1] = 2
	x[2] = 89
	x[3] = 76
	x[4] = 81
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", x)

	// slice
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s)

	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", s)

	// slicing a slice
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])

	// append
	s = append(s, 52)
	fmt.Println(s)

	s = append(s, 53, 54, 55)
	fmt.Println(s)

	t := []int{56, 57, 58, 59, 60}
	s = append(s, t...)
	fmt.Println(s)

	// delete
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	a = append(a[:3], a[6:]...)
	fmt.Println(a)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}

	y := []string{"James", "Bond", "Shaken", "not stirred"}
	z := []string{"Miss", "Moneypenny", "HElooooo", "James"}
	yz := [][]string{y, z}
	fmt.Println(yz)
	for _, v := range yz {
		fmt.Println(v)
		for _, w := range v {
			fmt.Println(w)
		}
	}

	m := map[string][]string{
		"bond_james":      []string{"shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"shaken, not stirred", "Martinis", "Women"},
		"no_dr":           []string{"shaken, not stirred", "Martinis", "Women"},
	}

	for k, val := range m {
		fmt.Println(k)
		for i, value := range val {
			fmt.Println("\t", i+1, value)
		}
	}

	m["fleming"] = []string{"cars"}
	for _, va := range m {
		fmt.Println(va)
	}

	delete(m, "fleming")
	for _, values := range m {
		fmt.Println(values)
	}
}
