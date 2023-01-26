package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	a := 0.1
	b := 0.2
	c := 0.3

	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, a+b == c)
	fmt.Printf("equal(%0.18f, %0.18f) : %v\n", c, a+b, equal(a+b, c))
}
