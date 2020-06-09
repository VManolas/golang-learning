package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {
	z := float64(1)
	counter := 0
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("%f\n", z)
		counter++
	}
	fmt.Printf("Counter is %d ", counter)
	return z
}

func Sqrt2(x float64) float64 {
	// z := float64(1)
	z := x
	counter := 0
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("%f\n", z)
		counter++
	}
	fmt.Printf("Counter is %d ", counter)
	return z
}

func Sqrt3(x float64) float64 {
	// z := float64(1)
	z := (x / 2)
	counter := 0
	for math.Abs(z*z-x) > 0.000001 {
		z -= (z*z - x) / (2 * z)
		//fmt.Printf("%f\n", z)
		counter++
	}
	fmt.Printf("Counter is %d ", counter)
	return z
}
func main() {
	x := 10.0
	for x <= 15 {
		fmt.Printf("Sqrt1(%f) = %.9f\n", x, Sqrt1(x))
		fmt.Printf("Sqrt2(%f) = %.9f\n", x, Sqrt2(x))
		fmt.Printf("Sqrt2(%f) = %.9f\n", x, Sqrt3(x))
		fmt.Printf("         math.Sqrt(%f) = %.9f\n", x, math.Sqrt(x))
		x++
	}

	y := 12.345
	fmt.Printf("%6.3f - %.3g\n", y, y)
}
