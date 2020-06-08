// Example 17a Methods
package main

import (
	"fmt"
	"math"
)

//  Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

//  Remember: a method is just a function with a receiver argument.
// Here's Abs written as a regular function with no change in functionality.
func funcAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// You can declare a method on non-struct types, too.
// You can only declare a method with a receiver whose type is defined in the same package as the method.
// You cannot declare a method with a receiver whose type is defined in another package
// (which includes the built-in types such as int)!!!
type MyFloat float32

func (f MyFloat) AbsMyFloat() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//  You can declare methods with pointer receivers.
// This means the receiver type has the literal syntax *T for some type T.
// (Also, T cannot itself be a pointer such as *int.)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Methods with pointer receivers can modify the value to which the receiver points (as Scale does here).
// Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

func funcScale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3.4, 4.6}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs()) //  methods with value receivers take either a value or a pointer as the receiver when they are called
	funcScale(&v, 10)
	fmt.Println(funcAbs(v))

	fmt.Println()
	p := &Vertex{3.4, 4.6}
	p.Scale(10)
	fmt.Println(p.Abs()) //  methods with value receivers take either a value or a pointer as the receiver when they are called
	// There are two reasons to use a pointer receiver:
	// 1. The first is so that the method can modify the value that its receiver points to.
	// 2. The second is to avoid copying the value on each method call. (This can be more efficient if the receiver is a large struct, for example.)
	// In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
	funcScale(p, 10)
	fmt.Println(funcAbs(*p))

	fmt.Println()
	f := MyFloat(-math.Sqrt2)
	fmt.Println(-f)
	fmt.Println(f.AbsMyFloat())
}
