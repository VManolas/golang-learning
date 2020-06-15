package main

import (
	"fmt"
	slt "golang-learning/Codility/Example02_WeekDay/solution"
	"math/rand"
)

func main() {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	K := rand.Intn(501)

	S := days[K%7]
	fmt.Printf("Random Day S is %s, and number of K days is %d (K mod 7 = %d days) \n", S, K, K%7)
	fmt.Println(slt.Solution(days, S, K))
}
