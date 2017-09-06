package main

import (
	"fmt"
	"strconv"
)

func addStringA(a int) string {
	i := strconv.Itoa(a)
	return i
}

func addStringB(b int) string {
	i := strconv.Itoa(b)
	return i
}

func add(x, y int) int {
	b := x + y
	return b
}

func main() {
	fmt.Println("THE RESULT OF", addStringA(42), "+", addStringB(13), "IS", add(42, 13))
}

// A function can take zero or more arguments.

// In this example, add takes two parameters of type int.

// Notice that the type comes after the variable name.

// (For more about why types look the way they do, see the article on Go's declaration syntax.)
// https://blog.golang.org/gos-declaration-syntax
