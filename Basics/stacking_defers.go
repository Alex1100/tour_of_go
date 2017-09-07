package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 10; 0 < i; i-- {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

// To learn more about defer statements read this blog post. https://blog.golang.org/defer-panic-and-recover
