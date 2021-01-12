package main

import "fmt"

// This is a package variable not a global variable
var (
	name      = "Mike"
	firstName = "Michael"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(name)
	fmt.Println(firstName)
}
