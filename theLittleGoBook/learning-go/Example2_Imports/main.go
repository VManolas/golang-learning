// Example 2a
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("It's over", os.Args[1])
	fmt.Println(os.Args[0])
}

// ---------------------------------------
// Example 2b
// Command: `go run main.go`
// Error message: ???

// Command: `go build main.go`
// Command: `./main`
// Error message: ???

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// }
