package main

import "fmt"

type ninjaSword struct {
	owner string
}

func (ninjaSword) attack() {
	fmt.Println("Swinging Ninja Sword")
}

type ninjaStar struct {
	owner string
}

func (ninjaStar) attack() {
	fmt.Println("Throwing Ninja Stars")
}

func (ninjaStar) pickUp() {
	fmt.Println("Pickup the Ninja Stars")
}

type ninjaWeapon interface {
	attack()
}

func main() {
	weapons := []ninjaWeapon{ninjaStar{"Wallace"}, ninjaSword{"Alex"}, ninjaStar{"Adam"}}

	for _, weapon := range weapons {
		weapon.attack()
		value, ok := weapon.(ninjaStar)
		if ok {
			value.pickUp()
		}

		switch value := weapon.(type) {
		case ninjaStar:
			value.pickUp()
		case ninjaSword:
			fmt.Println("No need to pickup again")
		}
	}
}
