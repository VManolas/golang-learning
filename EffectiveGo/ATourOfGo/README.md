## [A Tour of Go](https://tour.golang.org/)

### Examples
- [Example 1 Faking Time](Examples/Example01_FakingTime/main.go)
- [Example 2 Faking File System](Examples/Example02_FakingFileSystem/main.go)
- [Example 3 Faking The Network](Examples/Example03_FakingTheNetwork/main.go)
- [Example 4 Basic Type](Examples/Example04_BasicTypes/main.go)
- [Example 5 Type Conversions](Examples/Example05_TypeConversions/main.go)
- [Example 6 Numeric Constants](Examples/Example06_NumericConstants/main.go)
- [Example 7 For](Examples/Example07_For/main.go)
- [Example 8 If](Examples/Example08_If/main.go)
- [Example 9 Exercise Loops 'n' Functions](Examples/Example09_ExerciseLoopsFunctions/main.go)
- [Example 10 Switch](Examples/Example10_Switch/main.go)
- [Example 11 Defer](Examples/Example11_Defer/main.go)
- [Example 12 Pointers Struct Arrays](Examples/Example12_PointersStructArrays/main.go)
- [Example 13 Slices](Examples/Example13_Slices/main.go) (see also [relevant post on slices on golang's blog](https://blog.golang.org/slices). For more slice tricks check [here](https://github.com/golang/go/wiki/SliceTricks))
- [Example 14 Maps](Examples/Example14_Maps/main.go) (see also [relevant post on maps on golang's blog](https://blog.golang.org/maps)
)
- [Example 15 Function Values Closures](Examples/Example15_FunctionValuesClosures/main.go)
- [Example 16 Exercise Closures](Examples/Example16_ExerciseClosures/main.go)
- [Example 17 Methods](Examples/Example17_Methods/main.go)
- [Example 18 Interfaces](Examples/Example18_Interfaces/main.go)
- [Example 19 Type Assertions](Examples/Example19_TypeAssertions/main.go)
- [Example 20 Stringers](Examples/Example20_Stringers/main.go)
- [Example 21 Errors](Examples/Example21_Errors/main.go)
- [Example 22 Readers](Examples/Example22_Readers/main.go)
- [Example 23 Images](Examples/Example23_Images/main.go)

### Notes & Questions
- [Stateful vs. Stateless Architecture Overview](https://www.bizety.com/2018/08/21/stateful-vs-stateless-architecture-overview/)
    - [Defining statefull vs stateless web services](https://nordicapis.com/defining-stateful-vs-stateless-web-services/)
    - TL;DR,TR (*ToRead*) [Chapter 5 - Representational State Transfer (REST)](https://www.ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
- Packages: By convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand. 
- Exported names: In Go, a name is exported if it begins with a capital letter. For example, Pizza is an exported name, as is Pi, which is exported from the math package. 
- Functions: TR [Go's declaration syntax article](https://blog.golang.org/declaration-syntax)
  - "Naked" return: A return statement without arguments returns the named return values of the function
- Basic Types
  - [Strings, bytes, runes and characters in Go (from `The Go Blog`)](https://blog.golang.org/strings)
    - `rune`: _alias for int32 - represents a Unicode code point_
  - [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)
  - `uintptr` ?
- TR [Defer, Panic, and Recover](https://blog.golang.org/defer-panic-and-recover)
- Slice literas
  -  ~~A slice literal is like an array literal without the length ?~~
  -  ~~If I drop its first two values, can I re-gain them ?~~
- Test cases: [package wc/func test](https://github.com/golang/tour/blob/master/wc/wc.go#L10) ?
   -  [strings.Fields](https://golang.org/src/strings/strings.go?s=8396:8426#L319)
- `Git`: `git add -u`, `git rebase --skip`, `git log -p`
- `Vim` shortcuts: {`$`:EOL, `dd`:cut, `p`:paste, `x`:remove char, `Ctrl+L`: clear bash screen

### Other Questions
1. [Go/Playground](https://blog.golang.org/playground): Example1 Faking Time ?
2. [Go/Playground](https://blog.golang.org/playground): Example2 Faking File system ?
3. [Go/Playground](https://blog.golang.org/playground): Example3 Faking the Network ?
4. [Note: Time in the Go playground](https://tour.golang.org/flowcontrol/10) always appears to start at 2009-11-10 23:00:00 UTC, a value whose significance is left as an exercise for the reader ([Example_10](Examples/Example10_Switch/main.go)) ?