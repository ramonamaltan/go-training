package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favIceCream []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	forWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := person{
		firstName:   "Ramona",
		lastName:    "Maltan",
		favIceCream: []string{"cookies", "chocolate"},
	}

	p2 := person{
		firstName:   "Anna",
		lastName:    "Wintour",
		favIceCream: []string{"yoghurt", "strawberry", "vanilla"},
	}

	fmt.Println(p1, p2)
	for i, v := range p2.favIceCream {
		fmt.Println(i+1, v)
	}

	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(m)

	for k, val := range m {
		fmt.Println("\n", k, val)
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "gold",
		},
		forWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(t1, s1)
	fmt.Println(t1.doors)
	fmt.Println(s1.luxury)

	m1 := struct {
		courses int
		starter []string
		main    string
		dessert map[int]string
	}{
		courses: 4,
		starter: []string{"pomodoro", "pasta"},
		main:    "pizza",
		dessert: map[int]string{
			1: "strawberry",
			2: "vanilla",
		},
	}

	fmt.Println(m1)
}
