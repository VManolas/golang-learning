// Example 10a
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	lookup := make(map[string]int)
// 	lookup["goku"] = 9001
// 	power, exists := lookup["vegeta"]
// 	fmt.Println(power, exists)
// 	power, exists = lookup["goku"]
// 	fmt.Println(power, exists)

// 	total := len(lookup)
// 	fmt.Println(total)
// 	fmt.Println(lookup)
// 	delete(lookup, "vegeta")
// 	fmt.Println(total)
// 	fmt.Println(lookup)
// 	delete(lookup, "goku")
// 	fmt.Println(total)
// 	fmt.Println(lookup)
// }

// Example 10b
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	lookup := make(map[string]int, 100) // initial size - can help with performance
// 	fmt.Println(lookup)
// }

// Example 10c
package main

import "fmt"

// when you need a map as a field of a structure
type Saiyan struct {
	Name    string
	Friends map[string]*Saiyan
}

// a way to initialize the above is via:
func main() {
	goku := &Saiyan{
		Name:    "Goku",
		Friends: make(map[string]*Saiyan),
	}
	goku.Friends["krillin"] = goku //load or create krillin ??

	lookup := map[string]int{
		"goku":  9001,
		"gohan": 2044,
		"konan": 20000,
	}
	for key, value := range lookup {
		fmt.Println(key, value) // iteration over maps isn't ordered. Each iteration over lookup will return the key value pair in a random order
	}
}
