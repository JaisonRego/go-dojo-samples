package main

import (
	"errors"
	"fmt"
)

func function() {
	fmt.Println("Inside Function")
}

func variadic(integers ...int) {
	for _, value := range integers {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func sum(a, b int) int {
	return a + b
}

func nameLength(name string) (string, int) {
	return name, len(name)
}

func greeting(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Empty Name! This is not good")
	} else {
		return "Hello " + name, nil
	}
}

func print(result string, err error) {
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func fibonacci(i int) int {
	switch i {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return fibonacci(i-1) + fibonacci(i-2)
	}
}

func pointer(i *int) {
	*i++
}

func returnParam(i int) int {
	return i
}

func returnFunction(i int, f func(int) int) int {
	fmt.Printf("Function type %T\n", f)
	return f(i)
}

func first() {
	fmt.Println("First")
}

func second() {
	fmt.Println("Second")
}

func main() {
	function()
	variadic(1, 2, 3, 4)

	fmt.Println(sum(2, 4))
	fmt.Println(nameLength("Wallace"))

	print(greeting(""))
	print(greeting("Wallace"))

	fmt.Println(fibonacci(12))

	i := 1
	fmt.Println(i)
	pointer(&i)
	fmt.Println(i)

	fmt.Printf("The type returning is %T\n", returnParam)
	fmt.Printf("The type returning is %T and value is %d\n", returnParam(1), returnParam(1))
	fmt.Printf("The value: %d\n", returnFunction(1, returnParam))

	defer first()
	second()

	aFunction := func() {
		fmt.Println("I am in aFunction")
	}
	aFunction()

	func() {
		fmt.Println("anonymous function")
	}()

	// crazishFunction := func(f func()) func() {
	// 	return f
	// }

	// crazyFunction(crazishFunction)(first)()
}

// func crazyFunction(f func(func()) func()) func(func()) func() {
// 	return f
// }
