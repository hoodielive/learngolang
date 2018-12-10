package main

import "fmt"

// package scope
var x int = 42 // initialize variable

func main() {
	fmt.Println(x)
	foo()
	y := 17
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
}
