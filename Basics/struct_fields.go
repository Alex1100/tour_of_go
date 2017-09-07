package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y float64
	Z string
	A int64
}

func main() {
	v := Vertex{1, -2.0, "hello", -33}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)
	fmt.Println(v.Z)
	fmt.Println(v.A)
}

// Struct Fields
// Struct fields are accessed using a dot.
