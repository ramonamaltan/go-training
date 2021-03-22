package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string
	Age   int
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	sort.Sort(ByAge(users))
	fmt.Println(users)
}
