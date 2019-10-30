package main

import "fmt"

func main() {
	var found bool
	carToLookFor := "Blazer"
	cars := []string{"Accord", "IS250", "Blazer"}
	for _, car := range cars {
		if car == carToLookFor {
			found = true // set flag
		}
	}
	fmt.Printf("Found? %v", found)
}
