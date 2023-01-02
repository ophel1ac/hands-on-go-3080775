// flow-control/switch/begin/main.go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS

	// refactor with switch statement
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "linux", "darwin", "unix":
		fmt.Println("*nix variant")
	default:
		fmt.Printf("%s.\n", os)
	}
}
