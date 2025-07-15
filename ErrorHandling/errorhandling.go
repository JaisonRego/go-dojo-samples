package main

import (
	"errors"
	"fmt"
	"strconv"
)

func returnErrors(returnError bool) (string, error) {
	if returnError {
		return "", errors.New("error here")
	}
	return "Random Result!!", nil
}

type specialError struct {
	errorMessage string
	errorCode    int
}

func (s specialError) Error() string {
	return s.errorMessage + " with error code " + strconv.Itoa(s.errorCode)
}

func safeError() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic state!!! caused by ->", r)
	}
}

func main() {
	result, err := returnErrors(true)
	if err != nil {
		fmt.Println(specialError{"Not aunthorised", 1004})
	} else {
		fmt.Println(result)
	}

	const one = 2
	defer safeError()
	if one != 1 {
		panic("Value is not same")
	}

}
