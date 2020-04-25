package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, f, z)

	// Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
	var x1, y1 int = 2, 3
	//var f1 float64 = math.Sqrt((x1*x1 + y1*y1)) // cannot use x1 * x1 + y1 * y1 (type int) as type float64 in argument to math.Sqrt
	var f1 = math.Sqrt(float64(x1*x1 + y1*y1))
	//var z1 uint = f1 //cannot use f1 (type float64) as type uint in assignment
	var z1 uint = uint(f1)
	fmt.Println(x1, y1, f1, z1) // ??
}
