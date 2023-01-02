// flow-control/loop-range/begin/main.go
package main

import "fmt"

func main() {
	// initialize a slice of ints
	//
	s := []int{1, 2, 3, 4, 5}

	// use for-range to iterate over the slice and print each value
	//
	for i, val := range s {
		fmt.Printf("%v ", val)
		if i == len(s)-1 {
			fmt.Println()
		}
	}

	// declare a map of strings to ints
	//
	m := map[string]int{"Tears of Joy": 602, "Heart": 2764}

	// use for-range to iterate over the map and print each key/value pair
	//
	for key, val := range m {
		fmt.Printf("key=%v, value=%v\n", key, val)
	}
}
