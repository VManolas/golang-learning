// Example 16a Exercise: Fibonacci Closures
// Implement a fibonacci function that returns a function (a closure) that returns
// successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var i, j int
	return func() int {
		i, j = j, i+j
		if j == 0 {
			j = 1
		}
		return i
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f())
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
