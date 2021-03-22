package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Miss","Last":"Moneypenny","Age":27}]`
	bs := []byte(s)

	type person struct {
		First string
		Last  string
		Age   int
	}

	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("all data", people)

	for i, v := range people {
		fmt.Println("\nPERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
