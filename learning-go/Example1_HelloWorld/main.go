// Example 1a
package main

func main() {
	println("It's over 9000!")
}

// ---------------------------------------
// Example 1b
// Command: `go run main.go`
// Error message: `go run: cannot run non-main package`

// Command: `go build main.go`
// Command: `./main`
// Message: `It's over 9000!`

// package heap

// func main() {
// 	println("It's over 9000!")
// }

// ---------------------------------------
// Example 1c
// Command: `go run main.go`
// Error message: `runtime.main_main·f: function main is undeclared in the main package`
// Command: `go build main.go`
// Error message: `runtime.main_main·f: function main is undeclared in the main package`

// package main

// func example1() {
// 	println("It's over 9000!")
// }
