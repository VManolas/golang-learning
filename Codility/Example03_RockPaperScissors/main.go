package main

import (
	"fmt"
	slt "golang-learning/Codility/Example03_RockPaperScissors/solution"
	"math/rand"
	"time"
)

const letterBytes = "RPS"

// RandStringBytes returns a random string which comprises of characters included in the const letterBytes
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))] // for further reading check [this from stackoverflow](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go)
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	G := RandStringBytes(rand.Intn(500))
	fmt.Printf("For the string %s\n", G)
	fmt.Printf("Franco's score %d \n", slt.Solution(G))
}
