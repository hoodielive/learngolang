package main

import "fmt"

const avogadro float64 = 6.0221413e+23

const grams = 100.0

type amu float64

func (mass amu) float() float64 {
	return float64(mass)
}

type metalloid struct {
	name string
	number int32
	weight amu
}

var metalloids = []metalloid {
	metalloid {"Boron", 5, 10.8},
	metalloid {"Antimony", 54, 90.2},
	metalloid {"Polonium", 7, 18.8},
}

// finds # of moles
func moles(mass amu) float64 {
	return float64(mass) / grams
}
