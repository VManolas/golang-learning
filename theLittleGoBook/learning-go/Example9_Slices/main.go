// Example 9a
// package main

// import "fmt"

// func main() {
// 	// var scoresArray1 [10]int
// 	// scoresArray1[0] = 339
// 	// var scoresArray2 = [4]int{339, 9000, 9001, 33} // why not :=
// 	// fmt.Println(len(scoresArray1), len(scoresArray2))
// 	// for range scoresArray1 {
// 	// 	fmt.Println(len(scoresArray1))
// 	// }
// 	// for index, value := range scoresArray2 {
// 	// 	fmt.Println(len(scoresArray2), index, value)
// 	// }
// 	scores := []int{1, 4, 293, 4, 9}
// 	// init the slice with a length of 10, and a capacity of 10 (the size of the underlying array)
// 	scores2 := make([]int, 10) // we use make instead of new because there is more to creating a slice than just allocating the memory for the underlying array (which is what new does). We also have to initialize the slice.
// 	// init the slice with a length of 0, and a capacity of 10
// 	scores3 := make([]int, 0, 10) // ?? note that make and len are overloaded

// 	fmt.Println(len(scores), len(scores2), len(scores3))
// }

// Example 9b
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores[7] = 9033
// 	fmt.Println(scores)
// // panic: runtime error: index out of range [7] with length 0

// // goroutine 1 [running]:
// // main.main()
// //         /Users/vasileiosmanolas/code/go/src/learning-go/Example9_Slices/main.go:33 +0x4b
// // exit status 2
// }

// Example 9c
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = append(scores, 5)
// 	fmt.Println(scores) // prints  [5] (array)
// 	fmt.Println(scores[0]) // prints 5 (value)
// }

// Example 9d - re-slice the slice
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = scores[0:8]
// 	scores[7] = 9033
// 	fmt.Println(scores)
// }

// Example 9e
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 0, 10)
// 	scores = scores[0:10]
// 	scores[9] = 9033
// 	scores = append(scores, 5) // now, array cappacity will be 2x initial capacity
// 	fmt.Println(scores, cap(scores))
// }

// Example 9f
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 0, 5)
// 	c := cap(scores)
// 	fmt.Println(c)

// 	for i := 0; i < 25; i++ {
// 		scores = append(scores, i)
// 		if cap(scores) != c {
// 			c = cap(scores)
// 			fmt.Println(c)
// 		}
// 	}
// 	fmt.Println(scores)
// }

// Example 9g
// package main

// import "fmt"

// func main() {
// 	scores := make([]int, 5)
// 	scores = append(scores, 9033)
// 	fmt.Println(scores)
// }

// Example 9h - 4 common ways to initialize a slice
// package main

// import "fmt"

// type Person struct {
// 	Name string
// }

// type Saiyans struct {
// 	*Person //because we didn't give it an explicit name, we can implicitly access the fields and functions of the composed type
// 	Power   int
// }

// func main() {
// 	names := []string{"Leto", "Jessica", "Paul"} // when you know the values that you want in the array ahead of time
// 	checks := make([]bool, 10)                   // when you'll be writing into specific indexes of a slice. For example see func extractPowers() below
// 	var names []string                           // a nil slice that is used in conjuction with append, when the numbers of elements is unknown
// 	scores := make([]int, 0, 20) // this version lets us specify an initial capacity; useful if we have an idea of how many elements we'll need. Even when you know the size, append can be used. For example see func extractPowersAppend() below
// }

// func extractPowers(saiyans []*Saiyans) []int {
// 	powers := make([]int, len(saiyans))
// 	for index, saiyan := range saiyans {
// 		powers[index] = saiyan.Power
// 	}
// 	return powers
// }
// func extractPowersAppend(saiyans []*Saiyans) []int {
// 	powers := make([]int, 0, len(saiyans))
// 	for _, saiyan := range saiyans {
// 		powers = append(powers, saiyan.Power)
// 	}
// 	return powers
// }

// Example 9i
// package main

// import "fmt"

// func main() {
// 	scores := []int{1, 2, 3, 4, 5}
// 	slice := scores[2:4]
// 	slice[0] = 999
// 	fmt.Println(scores)
// 	scores = scores[:len(scores)-1]
// 	fmt.Println(scores)
// 	scores = scores[1:]
// 	fmt.Println(scores)
// }

// Example 9j
// package main

// import "fmt"

// func main() {
// 	scores := []int{1, 2, 3, 4, 5}
// 	scores = removeAtIndex(scores, 2)
// 	fmt.Println(scores)
// }

// func removeAtIndex(source []int, index int) []int {
// 	lastIndex := len(source) - 1
// 	//swap the last value and the value we want to remove
// 	source[index], source[lastIndex] = source[lastIndex], source[index]
// 	return source[:lastIndex]
// }

// Example 9k
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	scores := make([]int, 100)
	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)
	scores = scores[1:]
	worst := make([]int, 5)
	copy(worst, scores[:5])
	fmt.Println(scores)
	fmt.Println(worst)
	copy(worst[2:4], scores[:5])
	fmt.Println(worst)
	copy(worst, scores[:8])
	fmt.Println(worst)
}
