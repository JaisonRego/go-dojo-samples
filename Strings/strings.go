package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Jaison Rego"
	fmt.Println(s)
	fmt.Println(s[0])
	fmt.Printf("%c\n", s[0])
	fmt.Println(s[0:6])
	fmt.Println(s[:6])
	fmt.Println(s[6:])
	fmt.Println(len(s))

	fmt.Println("Hello \tWorld")
	fmt.Println("Hello \nWorld")
	fmt.Println("Hello \bWorld")

	str := "abc"
	slice := []byte(str)
	fmt.Println(str)
	fmt.Printf("%s\n", slice)
	slice[2] = 97
	fmt.Printf("%s\n", slice)
	str = string(slice)
	fmt.Println(str)

	var test string
	fmt.Println(len(test))

	var sb strings.Builder
	fmt.Println(sb.Len(), sb.Cap())
	sb.WriteString("Jaison")
	fmt.Println(sb.Len(), sb.Cap(), sb.String())
	sb.Grow(100)
	fmt.Println(sb.Len(), sb.Cap(), sb.String())
	sb.Reset()
	fmt.Println(sb.Len(), sb.Cap(), sb.String())
	sb.Write([]byte{65, 65, 97})
	fmt.Println(sb.Len(), sb.Cap(), sb.String())

	p1 := 123
	str2 := strconv.Itoa(p1)
	fmt.Printf("%d, %T, %s, %T\n", p1, p1, str2, str2)
	str3, _ := strconv.Atoi(str2)
	fmt.Printf("%d, %T\n", str3, str3)
}
