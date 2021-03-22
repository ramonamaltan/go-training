package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	data := []byte(`[{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]`)

	var users []user

	err := json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println(users)

	for i, u := range users {
		fmt.Println("Person", i+1)
		fmt.Println("\tName:", u.First, ", Age:", u.Age)
	}
}
