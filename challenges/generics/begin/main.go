// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

func print[T any](v T) {
	fmt.Println(v)
}

// non-generic print functions

func printString(s string) { fmt.Println(s) }

func printInt(i int) { fmt.Println(i) }

func printBool(b bool) { fmt.Println(b) }

// Part 2 sum function refactoring
type numeric interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

// sum sums a slice of any type
func sum[T numeric](numbers ...T) T {
	var result T
	for _, n := range numbers {
		result += n
	}
	return result
}

func main() {
	// call generic print functions
	print("Hello")
	print(42)
	print(true)

	// call generic sum function
	fmt.Println("result", sum(1, 2, 3))
	fmt.Println("result", sum(1.1, 2.5, 3.2))
}
