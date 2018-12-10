package main

import "fmt"

/*
closure & anonymous functions
a function without a name

func expression
assigning a func to a variable like javascript :)
*/
func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}
