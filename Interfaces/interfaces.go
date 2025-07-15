package main

import "fmt"

type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon) {
	weapon.attack()
}

type ninjaStar struct {
	owner string
}

func (person ninjaStar) attack() {
	fmt.Println(person.owner, " is Throwing Ninja Star")
}

type ninjaSword struct {
	owner string
}

func (person ninjaSword) attack() {
	fmt.Println(person.owner, "is Attacking with Ninja Sword")
}

func main() {
	stars := []ninjaStar{{"Wallace"}, {"Adam"}}
	for _, weapon := range stars {
		weapon.attack()
	}
	fmt.Println()

	swords := []ninjaSword{{"Wallace"}, {"Alex"}}
	for _, weapon := range swords {
		weapon.attack()
	}
	fmt.Println()

	weapons := []ninjaWeapon{ninjaStar{"Adam"}, ninjaSword{"Alex"}, ninjaStar{"Wallace"}}
	for _, weapon := range weapons {
		attack(weapon)
	}

}
