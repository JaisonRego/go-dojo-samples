package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct {
	Name   string `json:"full_name"`
	Weapon Weapon `json:"weapon"`
	Level  int
}

type Weapon struct {
	Name  string `json:"weapon_name"`
	Level int    `json:"weapon_level"`
}

func main() {
	fmt.Println("Hello")

	sFrom := `{"full_name": "Wallace", "weapon": {"weapon_name": "Ninja Star", "weapon_level": 1}, "level": 1}`
	fmt.Println(sFrom)

	var wallace Ninja
	fmt.Println(wallace)

	err := json.Unmarshal([]byte(sFrom), &wallace)
	if err != nil {
		fmt.Println("Sad Boi", err)
	}
	fmt.Println(wallace)

	sTo, err := json.Marshal(wallace)
	if err != nil {
		fmt.Println("Sad Boi")
	}
	fmt.Println(string(sTo))

}
