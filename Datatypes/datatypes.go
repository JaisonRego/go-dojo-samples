package main

import (
	"fmt"
	"math"
)

func greeting(str string) {
	fmt.Println(str)
}

func main() {

	var v int = 2
	var p *int = &v
	fmt.Println(p, *p)

	type mytype int
	var c mytype
	fmt.Println(c)
	fmt.Printf("\n%T\n", c)

	a := 1
	b := 3.41
	b = float64(a)
	fmt.Println(b)

	fmt.Println(math.Sqrt(25))

	fmt.Println(complex(1, 2))
	var complex complex128 = complex(22, 33)
	fmt.Println(complex)

	venues := []string{"Home", "Work", "School", "Bar", "Gym"}

	for _, venue := range venues {
		switch venue {
		case "Home":
			greeting("Mom, I'm hungry")
		case "Work", "School":
			greeting("Weather is great today")
		case "Gym":
			greeting("Hey i like your dress but it is a little too blue")
		default:
			greeting("Sup bois")
		}
	}
}
