package main

import "fmt"

func main() {
	// key-value store
	// allow super fast efficient look-up
	// unordered list (arrays and slices are ordered)
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 24,
	}
	fmt.Println(m)
	// access value by key
	fmt.Println(m["James"])
	// enter non existing key -> 0
	// comma okay idiom to check if value exists in map
	v, ok := m["Hanna"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("this value exists", v)
	}

	// add key-value pair to a map
	m["todd"] = 33
	fmt.Println(m)

	// range over map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete from map
	// delete(<map name>, key)
	delete(m, "todd")
	fmt.Println(m)
}
