package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//  A map maps keys to values.
// The zero value of a map is nil. A nil map has no keys, nor can keys be added.
// The make function returns a map of the given type, initialized and ready for use.

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)
	fmt.Println(m["Bell Labs"])
	fmt.Println("len:", len(m))

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
