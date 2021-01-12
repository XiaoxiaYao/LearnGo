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
}
