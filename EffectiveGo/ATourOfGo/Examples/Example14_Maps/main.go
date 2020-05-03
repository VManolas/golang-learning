package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

type Vertex struct {
	Lat, Long float64
}

//  A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
// The make function returns a map of the given type, initialized and ready for use.

var m map[string]Vertex

func makeMapVertex() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println("len:", len(m))
	// Map literals are like struct literals, but the keys are required.
	m["Google"] = Vertex{
		37.42202, -122.08408,
	}
	fmt.Println(m)
	var o = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	fmt.Println(o)
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var p = map[string]Vertex{
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fmt.Println(p)
}

func makeOtherMap() {
	n := make(map[string]int)
	n["k1"] = 7
	n["k2"] = 13
	fmt.Println("map:", n)
	fmt.Println("len:", len(n))
	_, prs := n["k2"]
	fmt.Println("prs:", prs)

	delete(n, "k2")
	fmt.Println("map:", n)
	_, prs = n["k2"]
	fmt.Println("prs:", prs)
}

func makeAnotherMap() {
	// Mutating maps
	var mm = make(map[string]int)
	// Insert or update an element in map
	mm["answer"] = 42
	fmt.Println("The value: ", mm["answer"])
	mm["answer"] = 48
	fmt.Println("The value: ", mm["answer"])
	// Retrieve an element
	var answer = mm["answer"]
	fmt.Println(answer)
	// Delete an element
	delete(mm, "Answer")
	fmt.Println("The value:", mm["Answer"])
	// Test that a key is present with a two-value assignment
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := mm["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func main() {
	makeMapVertex()

	makeOtherMap()

	makeAnotherMap()
	// Exercise Maps
	// WordCount should return a map of the counts of each “word” in the string s.
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for key, value := range strings.Fields(s) {
		fmt.Println(key, value)
		m[value]++ // This is case sensitive.
	}
	return m
}
