package main

import "fmt"

/* declare, assign, initiliaze */

var b string // this is declaration; b = "golang" is assignment; this would be declaring a variable to its zero value

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	// godoc.org %v is default value for a type under package "fmt"
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}
