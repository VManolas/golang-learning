## [A Tour of Go](https://tour.golang.org/)

### Notes
- [Stateful vs. Stateless Architecture Overview](https://www.bizety.com/2018/08/21/stateful-vs-stateless-architecture-overview/)
    - [Defining statefull vs stateless web services](https://nordicapis.com/defining-stateful-vs-stateless-web-services/)
    - TL;DR,TR (*ToRead*) [Chapter 5 - Representational State Transfer (REST)](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
- Packages: By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand. 
- Exported names: In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. 
- Functions: TR [Go's declaration syntax article](https://blog.golang.org/declaration-syntax)
  - "Naked" return: A return statement without arguments returns the named return values of the function
- Basic Types
  - `rune` // alias for int32 - represents a Unicode code point ?
  - `uintptr` ?
- TR [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
- Slice literas
  -  A slice literal is like an array literal without the length ?
  -  If I drop its first two values, can I re-gain them ?


### Other Questions
1. [Go/Playground](https://blog.golang.org/playground): Example1 Faking Time ?
2. [Go/Playground](https://blog.golang.org/playground): Example2 Faking File system ?
3. [Go/Playground](https://blog.golang.org/playground): Example3 Faking the Network ?
4. [Note: Time in the Go playground](https://tour.golang.org/flowcontrol/10) always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader ([Example_10](Examples/Example10_Switch/main.go)) ?