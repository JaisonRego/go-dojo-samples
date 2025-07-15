package main

import "fmt"

func main() {
	const (
		c1 = 1
		c2
		c3 int = iota
		c4
	)
	fmt.Println(c1, c2, c3, c4)

	const (
		_ = 1 << (10 * iota)
		kb
		mb
		gb
	)
	fmt.Println(kb, mb, gb)

	fmt.Println()

	slice1 := []int{0, 1, 2}
	slice1 = append(slice1, 3, 4, 5, 6)
	slice1 = slice1[0:8]
	fmt.Println(slice1)
}
