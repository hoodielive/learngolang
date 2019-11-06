package main

import (
	"fmt"
	"math"
	"stringutil"
)

var myNumber = 1.23

func main() {

	roundedUp := math.Ceil(myNumber)
	roundedDown := math.Floor(myNumber)
	fmt.Println(roundedUp, roundedDown)
	fmt.Println(stringutil.Reverse("hello"))

}
