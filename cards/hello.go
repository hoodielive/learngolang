package main

import "fmt"

func main() { 
	card := []string{"I will live", newCard()}
	fmt.Println(card)
}

func newCard() string {
	return "An item - that looks like some data structure "
}