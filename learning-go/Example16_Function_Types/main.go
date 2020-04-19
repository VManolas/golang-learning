// Example 16a
package main

import "fmt"

type Add func(a int, b int) int

func main() {
	fmt.Println(process(func(a int, b int) int {
		return a + b
	}))
}

func process(adder Add) int {
	return adder(1, 2)
}

// Using functions like this can help decouple code from specific implementations much like we achieve with interfaces.
