// Example 15a Function Values
// package main

// import (
// 	"fmt"
// 	"math"
// )

// //  Functions are values too. They can be passed around just like other values.
// //  Function values may be used as function arguments and return values.
// func compute(fn func(float64, float64) float64) float64 {
// 	return fn(3, 4)
// }

// func triangle(x, y float64) float64 {
// 	return math.Sqrt(x*x + y*y)
// }

// func main() {
// 	fmt.Println(triangle(5, 12))
// 	fmt.Println(compute(triangle))
// 	fmt.Println(compute(math.Pow))
// }

// Example 15b Function Closures
// Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
