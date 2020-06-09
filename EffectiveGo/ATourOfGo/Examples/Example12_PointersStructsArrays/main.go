package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	// pointers
	i, j := 42, 2701

	p := &i           // point to i
	fmt.Println(*p)   // read i through the pointer
	*p = 21           // set i through the pointer
	fmt.Println(i, p) // see the new value of i

	p = &j            // point to j
	*p = *p / 37      // divide j through the pointer
	fmt.Println(j, p) // see the new value of j
	// structs
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		pv = &Vertex{1, 2} // has type *Vertex
	)
	fmt.Println(v1, pv, v3, v2)
	// arrays
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	// slices
	// An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
	//  A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	// a[low : high]
	// This selects a half-open range which includes the first element, but excludes the last one.
	var s []int = primes[1:4]
	fmt.Println(s)
}
