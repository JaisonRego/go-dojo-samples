package main

import "fmt"

func main() {
	array := [...]int{1: 4, 5: 2}
	fmt.Println(array)

	array2d := [...][2]int{{1, 2}, {3, 4}}
	fmt.Println(array2d)

	m := make(map[string]string, 5)
	fmt.Println(m)
	m = map[string]string{"Jaison": "Rego"}
	fmt.Println(m)
	delete(m, "Jaison")
	fmt.Println(m)

	title, ok := m["Allen"]

	if ok {
		fmt.Printf("Found Allen and he is %s\n", title)
	} else {
		fmt.Printf("Could not find Allen\n")
	}
}
