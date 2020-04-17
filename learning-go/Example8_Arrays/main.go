// Example 8a
package main

import "fmt"

func main() {
	var scores10 [10]int
	scores10[0] = 339

	var scores4 = [4]int{339, 9000, 9001, 33} // why not :=

	fmt.Println(len(scores4), len(scores10))

	for index, value := range scores4 {
		fmt.Println(len(scores4), index, value)
	}
	for range scores4 {
		fmt.Println(len(scores4))
	}
	for index, value := range scores10 {
		fmt.Println(len(scores10), index, value)
	}
}
