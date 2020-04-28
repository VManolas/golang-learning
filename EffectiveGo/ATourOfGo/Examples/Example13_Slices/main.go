package main

import (
	b64 "encoding/base64"

	"fmt"
	"strings"

	"golang.org/x/tour/pic" // https://pkg.go.dev/golang.org/x/tour/pic?tab=doc
)

func main() {
	// slices()
	// sliceLiterals()
	// sliceDefaults()
	// sliceLenCap()
	// nilSlice()
	// makeSlice()
	// sliceOfSlices()
	// appendToSlice()
	// sliceRange()
	pic.Show(Pic)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func slices() {
	// Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	b := names[0:2]
	c := names[1:3]
	fmt.Println(b, c)

	c[0] = "XXX"
	fmt.Println(b, c)
	fmt.Println(names)
}

func sliceLiterals() {
	// Slice literals
	// A slice literal is like an array literal without the length.
	// This is an array literal:
	// [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it:
	// []bool{true, true, false}
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)
}

func sliceDefaults() {
	// Slice defaults
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.
	sd := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(sd)
	sd = sd[1:4]
	fmt.Println(sd)

	sd = sd[:2]
	fmt.Println(sd)

	sd = sd[1:]
	fmt.Println(sd)
}

func sliceLenCap() {
	//  A slice has both a length and a capacity.
	// The length of a slice is the number of elements it contains.
	// The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	// The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	// You can extend a slice's length by re-slicing it, provided it has sufficient capacity.
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	// Extend the slice beyond its capacity
	//s = s[:4]
	//printSlice(s)
}

func nilSlice() {
	// Nil slices
	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var ns []int
	fmt.Println(ns, len(ns), cap(ns))
	if ns == nil {
		fmt.Println("nil!")
	}
}

func makeSlice() {
	// Creating a slice with make
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	// The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)
	printSlice(a)
	// To specify a capacity, pass a third argument to make:
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice(b)
	c := b[:cap(b)]
	printSlice(c)
	d := c[1:]
	printSlice(d)
}

func sliceOfSlices() {
	// Slices of slices
	// Slices can contain any type, including other slices.
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func appendToSlice() {
	// Appending to a slice
	// It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func sliceRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	//  You can skip the index or value by assigning to _.
	// for i, _ := range pow
	// for _, value := range pow
	// If you only want the index, you can omit the second variable.
	// for i := range pow
	pow2 := make([]int, 10)
	for i := range pow2 {
		pow2[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow2 {
		fmt.Printf("%d\n", value)
	}
}

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			//pic[i][j] = uint8((i+j) / 2)
			//pic[i][j] = uint8(i*j)
			pic[i][j] = uint8(i ^ j)
		}
	}
	// base64 decode function
	// import b64 "encoding/base64"
	data := "YmFzZTY0IGVuY29kZWQgc3RyaW5n"
	sDec, err := b64.StdEncoding.DecodeString(data)
	fmt.Println(sDec, err)
	return pic
}
