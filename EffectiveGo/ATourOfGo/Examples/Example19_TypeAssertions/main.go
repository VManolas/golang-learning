// Example 19a Type Assertions
// A type assertion provides access to an interface value's underlying concrete value.
package main

import (
	"fmt"
)

func main() {
	var i interface{} = "Hello"

	// t := i.(T)
	// This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

	s := i.(string)
	fmt.Println(s)
	// If i does not hold a T, the statement will trigger a panic.
	// f := i.(float64) // panic: interface conversion: interface {} is string, not float64
	// fmt.Println(f)

	// t, ok := i.(T)
	// To test whether an interface value holds a specific type, a type assertion can return two values:
	// the underlying value and a boolean value that reports whether the assertion succeeded.
	//  If i holds a T (string), then s will be the underlying value and ok will be true.
	s, ok := i.(string)
	fmt.Println(s, ok)
	// If not, ok will be false and f will be the zero value of type T (float64), and no panic occurs.
	f, ok := i.(float64)
	fmt.Println(f, ok)
}

// Example 19b Type switches
// A type switch is a construct that permits several type assertions in series.
// A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
// and those values are compared against the type of the value held by the given interface value.
// package main

// import (
// 	"fmt"
// )

// func do(i interface{}) {
//	// The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.
// 	switch v := i.(type) {
// 	case int:
// 		fmt.Printf("Twice %v is %v\n", v, v*2)
// 	case string:
// 		fmt.Printf("%q is %v bytes long\n", v, len(v))
// 	default:
// 		fmt.Printf("I don't know about type %T!\n", v)
// 	}
// }

// func main() {
// 	do(21)
// 	do("Hi")
// 	do(true)
// 	do(3.14)
// }
