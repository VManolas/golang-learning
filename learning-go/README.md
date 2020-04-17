This is the README file for my Learning Go attempt.
It has the format of a "diary".
If you are also starting off, you may find these resources useful.

### Notes
- [The Little Go Book Repository (Latest Version)](https://github.com/karlseguin/the-little-go-book)
  - [Effective Go documentation](https://golang.org/doc/effective_go.html)
  - [Golang-nuts Google Group](https://groups.google.com/forum/#!forum/golang-nuts) for when you are seeking for help
  - [Go playground](https://play.golang.org/) lets you run code online
  - [Installing Go](https://golang.org/dl/)
  - [Go Packages Documentation](https://golang.org/pkg/)
- Go is a compiled, statically typed language with a C-like syntax and garbage collection
- Structures, methods, pointers, functions on structures
  - _There is similarity between the definition of a structure and that of a class_
-  Constructors: **structures don't have constructors**. Instead, you create a function that returns and instance of the desired type (like a factory) 
- Composition
- Pointers VS Values
- Arrays
  - In Go, arrays are fixed
- Slices
  - A slice is a lightweight structure that wraps and represents a portion of an array
  - `Append`: if the underlying array is full, it will create a new larger array and copy the values over. This is why we re-assign the value returned
  - Slices as wrappers to arrays is a powerful concept
- Maps (key: value)
  - Are created with the `make` function (like slices)
  - Grow dynamically
- Packages 
  - when you name a package, via the **package** keyword, you provide a single value, not a complete hierarchy (e.g. "shopping" or "db")
  - when you import a package, you specify the complete path

### Examples
- Example 1 Hello World
    - `go run main.go` : uses a temp directory to build the program, executes it and then cleans itself up. Check the location of the temp file `go run --work main.go`
    - `go build main.go` : generates an executable binary `main` (`./main`)
- Example 2 Imports
- Example 3 Variables & Declarations
- Example 4 Function Declarations
  - _ is the blank identifier
- Example 5 Structures
- Example 6 Constructors 
- Example 7 Composition
  - Go supports composition: the act of including one structure into another. 
  - Is Composition better than Inheritance? Many people think that it's a more robust way to share code. When using Inheritance, your class is tightly coupled to your superclass and you end up focusing on hierarchy rather than behavior.
  - Overloading: Go doesn't support overloading.
- Example 8 Arrays
- Example 9 Slices
- Example 10 Maps
- Example 11 Packages
  - See directory `../shopping`

### Questions
- What are the ENV variables in ~/.profile ?
- Why `Example1/main.go/Example 1b` prints the correct result when using `go build main.go`, `./main` commands ?
  - Notice that the code compiles, there's just no entry point to run it. This is perfectly normal when you are, for example, building a library (?)
- Might have been upgraded but couldn't use `godoc -http=:6000` (localhost:6000)
- How to check which editor I use as default in vscode?
- In a terminal, I can type "vi(m)" to open Vi(m) editor. I can also type "bash" to open Bash editor. Which one is the default? 
- Would it make sense to start using vim editor straight away?
- What is composition and what is the "battle cry" about it over inheritance (OOP)?
- Functions on structures: 
  - func (s *Saiyan) Super() {	s.Power += 10000  } VS func Super(s *Saiyan) {  s.Power += 10000  } ?
- What is the factory pattern (example 6)? And How it shields the rest of the code from worrying about allocation details?
- See ?? in Example7/main.go
- Example 9h - func extractPowers
- 