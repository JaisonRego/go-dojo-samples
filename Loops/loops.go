package main

import "fmt"

func main() {

	for i, j := 2, 1; i < 10 || j < 6; i, j = i+2, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println()

	s := "Hello World"
	for _, c := range s {
		if c == ' ' {
			break
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	for _, c := range s {
		if c == ' ' {
			continue
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				break outerLoop
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}
