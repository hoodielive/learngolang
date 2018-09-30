package main

import "fmt"

func main() {
	const x string = "This is a constant of type string"
	fmt.Println(x)

	x = "Should get an error saying: cannot assign x"
	fmt.Println(x)
}
