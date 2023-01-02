// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error while recovering", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	//
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	// convert the bytes to a string
	fileData := string(data)

	// initialize a map to store the counts
	counts := make(map[string]int)

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(fileData)

	// capture the length of the words slice
	counts["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, char := range word {
			switch {
			case unicode.IsDigit(char):
				counts["numbers"]++
			case unicode.IsLetter(char):
				counts["letters"]++
			default:
				counts["symbols"]++ 
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(counts)
}
