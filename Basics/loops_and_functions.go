package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	min_delta := 0.00000000001
	delta := 1.0
	i := 0
	for ; math.Abs(delta) > min_delta; i++ {
		delta = (z*z - x) / (2 * z)
		z = z - delta
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

// As a simple way to play with functions and loops, implement the square root function using Newton's method.

// In this case, Newton's method is to approximate Sqrt(x) by picking a starting point z and then repeating:

// To begin with, just repeat that calculation 10 times and see how close you get to the answer for various values (1, 2, 3, ...).

// Next, change the loop condition to stop once the value has stopped changing (or only changes by a very small delta). See if that's more or fewer iterations. How close are you to the math.Sqrt?

// Hint: to declare and initialize a floating point value, give it floating point syntax or use a conversion:

// z := float64(1)
// z := 1.0
