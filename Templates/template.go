package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name   string
	Secret string
}

func error_log(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Hello")

	secret := secret{"Wallace", "I'm a Ninja"}
	templatePath := "/Users/jaisonrego/go/src/templates/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)
	error_log(err)

	err = t.Execute(os.Stdout, secret)
	error_log(err)
}
