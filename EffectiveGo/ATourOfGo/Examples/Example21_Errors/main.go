// Example 21a Errors
// Go programs express error state with error values.
// The error type is a built-in interface similar to fmt.Stringer:
// type error interface {
//   Error() string
// }
// Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.
// A nil error denotes success; a non-nil error denotes failure.
// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"time"
// )

// type MyError struct {
// 	When time.Time
// 	What string
// }

// func (e *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s", e.When, e.What)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"It didn't work mate, sorry!",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}

// 	i, err := strconv.Atoi("42E")
// 	if err != nil {
// 		fmt.Printf("couldn't convert number: %v\n", err)
// 		return
// 	}
// 	fmt.Println("Converted integer: ", i)
// }

// Example 21b Exercise Errors
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %g\n", float64(e))
}

const e = 1e-8 // small delta

func Sqrt(x float64) (float64, string) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x).Error()
	}
	// starting point
	z := x
	for {
		new_z := z - ((z*z - x) / (2 * z))
		if math.Abs(new_z-z) < e {
			return new_z, "Bingo!"
		}
		z = new_z
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2.12345678901234567890))
}
