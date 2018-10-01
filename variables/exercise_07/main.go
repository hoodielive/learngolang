package main

import (
	"strutils"
	"fmt"
)

func main() {
	name := "larry"
	strutils.ToUpperCase(name)
	fmt.Println(name)
}
