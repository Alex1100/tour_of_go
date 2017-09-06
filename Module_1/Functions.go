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
