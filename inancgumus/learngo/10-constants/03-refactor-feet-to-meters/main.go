// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"os"
	"strconv"
)

// EXERCISE
//
// This program uses magic numbers
// Refactor it and use named constants instead

func main() {
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048
	yards := feet * 0.3333

	fmt.Printf("%g feet is %g meters.\n", feet, meters)
	fmt.Printf("%g feet is %g yards.\n", feet, yards)
}
