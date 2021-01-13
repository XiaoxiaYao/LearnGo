package main

import "fmt"

// This is a package variable not a global variable
var (
	name      = "Mike"
	firstName = "Michael"
)

func main() {
	const PI = 3.14
	const (
		cpp = iota
		java
	)
	fmt.Println("Hello world")
	fmt.Println(name)
	fmt.Println(firstName)
	fmt.Println(PI)
	fmt.Println(java)

	if PI == 3.14 {
		fmt.Println("It's equal to 3.14")
	}

	var result int
	switch result {
	case 0:
		fmt.Println("Hi")
	default:
		panic("unsupported")
	}

	fmt.Println(grade(0))

	array1 := [...]int{1, 2}
	fmt.Println(array1)
	for index, item := range array1 {
		fmt.Println(index, item)
	}

	s2 := make([]int, 16)
	fmt.Println(s2)

	m := map[string]string{
		"name": "michael",
	}
	fmt.Println(m)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	default:
		panic(fmt.Sprint("wrong: %d", score))
	}

	return g
}
