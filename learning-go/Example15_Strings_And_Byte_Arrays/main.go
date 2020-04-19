// Example 15a
package main

import "fmt"

func main() {
	stra := "the spice must flow"
	byts := []byte(stra)
	strb := string(byts)

	fmt.Println(stra)
	fmt.Println(byts)
	fmt.Println(strb)

	fmt.Println(len("￿"))
	// fmt.Println(range byts) ??
}
