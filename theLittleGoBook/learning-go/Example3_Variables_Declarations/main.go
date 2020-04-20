// Example 3a
package main

import (
	"fmt"
)

func main() {
	var power int
	power = 9000
	fmt.Printf("It's over %d\n", power)
}

// ---------------------------------------
// Example 3b
// Command: `go run --work main.go`

// Command: `go build main.go`
// Command: `./main`

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var power int = 9000
// 	fmt.Printf("It's over %d\n", power)
// }

// ---------------------------------------
// Example 3c
// Command: `go run --work main.go`

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	power := 9000
// 	fmt.Printf("It's over %d\n", power)
// }

// ---------------------------------------
// Example 3d
// Command: `go run --work main.go`

// package main

// import "fmt"

// func main() {
// 	power := getPower()
// 	fmt.Printf("It's over %d\n", power)
// }

// func getPower() int {
// 	return 9001
// }

// ---------------------------------------
// Example 3e
// Command: `go run --work main.go`
// Error message: "no new variables on left side of :="
// package main

// import "fmt"

// func main() {
// 	power := getPower()
// 	fmt.Printf("It's over %d\n", power)

// 	power := power + 2
// 	fmt.Printf("It's over %d\n", power)
// }

// func getPower() int {
// 	return 9001
// }

// ---------------------------------------
// Example 3f
// Command: `go run --work main.go`

// package main

// import "fmt"

// func main() {
// 	name, power := "Goku", getPower()
// 	fmt.Printf("%s's power is over %d\n", name, power)
// }

// func getPower() int {
// 	return 9001
// }

// ---------------------------------------
// Example 3g
// Command: `go run --work main.go`

// package main

// import "fmt"

// func main() {
// 	power := getDefaultPower()
// 	fmt.Printf("Default power is %d\n", power)

// 	name, power := "Goku", getUserPower()
// 	fmt.Printf("%s's power is %d\n", name, power)
// }

// func getDefaultPower() int {
// 	return 1000
// }

// func getUserPower() int {
// 	return 9001
// }

// ---------------------------------------
// Example 3h
// Command: `go run --work main.go`

// package main

// import "fmt"

// func main() {
// 	// Error Message:
// 	// # command-line-arguments
// 	// ./main.go:134:2: name declared but not used
// 	name, power := "Goku", getDefaultPower()
// 	fmt.Printf("Default power is %d\n", power)
// }

// func getDefaultPower() int {
// 	return 1000
// }
