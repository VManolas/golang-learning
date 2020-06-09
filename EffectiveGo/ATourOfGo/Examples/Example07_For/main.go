package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	// 	for ; sum < 1000; { // same as
	for sum < 1000 { // which is the while(){} of C
		sum += sum
	}
	fmt.Println(sum)

	//infinite loop
	for {
		sum += sum
	}
}
