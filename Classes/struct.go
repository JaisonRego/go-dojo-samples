package main

import "fmt"

type ninjaIntern struct {
	name  string
	level int
}

func (n ninjaIntern) sayHi() {
	fmt.Printf("Hello, welcome the Dojo %s\n", n.name)
}

func main() {
	type ninja struct {
		name    string
		weapons []string
		level   int
	}

	wallace := ninja{name: "Wallace"}
	fmt.Println(wallace)
	wallace = ninja{
		name:    "jaison",
		weapons: []string{"None"},
		level:   2,
	}
	fmt.Println(wallace)
	wallace = ninja{"Wallace", []string{"Ninja Star"}, 3}
	fmt.Println(wallace)
	wallace.level++
	wallace.weapons = append(wallace.weapons, "Ninja Sword")
	fmt.Println(wallace)

	type dojo struct {
		name  string
		ninja ninja
	}

	dojo1 := dojo{
		name:  "MyDojo",
		ninja: wallace,
	}
	fmt.Println(dojo1)
	dojo1.ninja.level = 5
	fmt.Println(dojo1)
	fmt.Println(wallace)

	type addressDojo struct {
		name  string
		ninja *ninja
	}

	address := addressDojo{
		name:  "Share Dojo",
		ninja: &wallace,
	}

	fmt.Println(address)
	fmt.Println(*address.ninja)
	address.ninja.level = 5
	fmt.Println(*address.ninja)
	fmt.Println(wallace)
}
