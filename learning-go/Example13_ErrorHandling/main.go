// Example 13a
// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// // we can create our own errors by importing the `errors` package and using it in the `New` function
// func process(count int) error {
// 	if count < 1 {
// 		return errors.New("Invalid Count")
// 	}
// 	// ...
// 	return nil
// }

// func main() {
// 	if len(os.Args) != 2 {
// 		os.Exit(1)
// 	}

// 	n, err := strconv.Atoi(os.Args[1])
// 	if err != nil {
// 		fmt.Println("not a valid number")
// 	} else {
// 		fmt.Println(n)
// 	}
// }

// Example 13b - singleton
package main

import (
	"fmt"
	"io"
)

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err == io.EOF {
		fmt.Println("no more input!")
	}
}
