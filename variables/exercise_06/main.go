package main

import "fmt"

func main() {
	
	var fname string 

	var lname string 
	
	var age int
	 
	fname = "rufus"
	lname = "parrot"
	age = 98

	fmt.Println("First Name:", fname)

	fmt.Println("Last Name:", lname)


	fmt.Println("Age:", age)

	addFunc()
}


func addFunc() {

	var fname,lname string = "Shiju", "Veritable"
	fmt.Println(fname, lname)
}

