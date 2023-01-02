// flow-control/loop-basic/begin/main.go
package main

import "fmt"

func main() {
	// declare a string to iterate over
	//
	str := "some string"

	// iterate over the string with basic for loop
	//
	for _, l := range str {
		fmt.Printf("%s", string(l))
	}
}
