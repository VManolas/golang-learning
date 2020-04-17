// Example 4a
package main

import "fmt"

func main() {
	// Error message: value declared but not used
	//value, exists := power("goku")
	// _ is the blank identifier
	// the return value is not actually assigned which allows to use _ over and over again regardless of the returned type
	_, exists := power("goku")
	if exists == false {
		fmt.Printf("Power not exists\n")
	}
}

//func log(message string)   {}
//func add(a int, b int) int {}
// If parameters share the same type
// func add(a, b int) int {}
func power(name string) (int, bool) {
	return 1, false
}
