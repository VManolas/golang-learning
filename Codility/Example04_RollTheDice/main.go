package main

import (
	"fmt"
	slt "golang-learning/Codility/Example04_RollTheDice/solution"
)

func main() {
	A := [...]int{3, 2, 4, 3} // The N roll results that you remember are described by an array A
	F := 3                    // Rolls whose results you have forgotten
	M := 4                    // The arithmetic mean of all the roll results (the sum of all the roll results divided by the number of rolls)
	fmt.Println(slt.Solution(A[:], F, M))
}
