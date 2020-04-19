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
  - Package Management: `go get`
    - `go get` **within a project**, it’ll scan all the files, looking for imports to third-party libraries and will download them. In a way, **our own source code becomes a Gemfile or package.json.**
    - `go get -u` updates the packages (`go get -u FULL_PACKAGE_NAME` updates a specific package)
    - There is no way to specify a revision, it always points to the `master/head/trunk/default`. This is an even larger problem if you have two projects needing different versions of the same library. To solve this you can use a [3rd-party dependency management tool](https://github.com/golang/go/wiki/PackageManagementTools) (e.g. `goop`, `godep`)
- Interfaces
  - define **a contract** but **not an implementation**
  - (what purpose?) help **decouple code from specific implementations**. (How?) We might have various types of loggers, yet by programming against the interface, rather than these concrete implementations, we can easily change (and test) which we use without any impact to our code.
  - In a language like C# or Java, we have to be **explicit** when a class implements an interface. In Go, this happens **implicitly**. This cuts down on the verboseness of using interfaces, and also ends to promote small and focused interfaces. 
  - The standard library is full of interfaces. The `io` package has a handful of popular ones such as `io.Reader`,`io.Writer`, and `io.Closer`. If you write a function that expects a parameter that you’ll only be calling `Close()` on, you absolutely should accept an `io.Closer` rather than whatever concrete type you’re using.
  - Interfaces can also participate in **composition**. And, interfaces themselves can be composed of other interfaces. For example, `io.ReadCloser` is an interface **composed** of the `io.Reader` interface as well as the `io.Closer` interface.
  - Interfaces are commonly used to **avoid cyclical imports**. Since they don’t have implementations, they’ll have **limited dependencies**.
- Tidbits
  - Error handling
    - Go’s preferred way to deal with errors is through **return values**, **not exceptions**. 
    - You can create your own error type; the only requirement is that it fulfills the contract of the built-in `error` interface, which is 
    ```
    type error interface {
      Error() string
    }
    ```
    - There is a common pattern in Go's standard library of using error variables. For example, the `io` package has an `EOF` variable which is defined as: 
    ```
    var EOF = errors.New("EOF")
    ```
    - This is a package variable (it's defined outside of a function) which is publicly accessible (upper-case first letter). Various functions can return this error, say when we're reading from a file or STDIN. If it makes contextual sense, you should use this error, too.
    - Go does have `panic` and `recover` functions. `panic` is like throwing an exception while `recover` is like `catch`; they are rarely used. 
  - Defer
    - Even though Go has a garbage collector, **some resources require that we explicitly release them**. For example, we need to `Close()` **files** after we’re done with them. This sort of code is always dangerous. For one thing, as we’re writing a function, it’s easy to forget to Close something that we declared 10 lines up. For another, a function might have multiple return points. Go’s solution is the `defer` keyword.
    - Whatever you defer will be **executed after the enclosing function returns**, even if it does so violently. This lets you release resources near where it’s initialized and takes care of multiple return points.
  - `go fmt`
    - When you’re inside a project, you can apply the formatting rule to it and all sub-projects via: `go fmt ./...`
  - Initialized if
    - a value can be initiated prior to the condition being evaluated. While the values aren’t available outside the if-statement, they are available inside any `elseif` or `else` 
    ```
    if err := process(); err!=nil {
      return err
    }
    ```
  - Empty Interfaces and Conversions
    - In most object-oriented languages, a built-in base class, often named `object`, is the **superclass** for all other classes. Go, *having no inheritance, doesn’t have such a superclass*. What it does have is **an empty interface with no methods**: `interface{}`. Since every type implements all 0 of the empty interface’s methods, and since interfaces are implicitly implemented, every type fulfills the contract of the empty interface.
    - If we wanted, we could write an `add` function with the following signature:
    ```
    func add(a interface{}, b interface{}) interface{} { ... }
    ```
    - To convert an interface variable to an explicit type, you use `.(TYPE)`:
    ```
    return a.(int) + b.(int) 
    // if the underlying type is not `int`, the above will result in an error
    ```
    - A powerful **type switch**:
    ```
    switch a.(type) {
      case int:
        fmt.Printf("a is now an int and equals %d\n", a)
      case bool,string: 
        //...
      default:
        //...
    }
    ```
    - *You’ll see and probably use the empty interface more than you might first expect. Admittedly, it won’t result in clean code. Converting values back and forth is ugly and dangerous but sometimes, in a static language, it’s the only choice*.
  - Strings and Byte Arrays
    - when you use `[]byte(X)` or `string(X)`, you’re creating a copy of the data. This is necessary because **strings are immutable**.
    - Strings are made of `runes` which are *unicode code points*. If you take the length of a string, you might not get what you expect.
    - If you iterate over a string using `range`, you’ll get `runes`, not bytes. Of course, when you turn a string into a `[]byte` you’ll get the correct data.
  - Function Type
    - Functions are 1st class types (`type Add func (a int, b int) int`) which can be used anywhere - as a field type, a parameter, a return value.
- Concurrency

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
- Example 12 Interfaces
- Example 13 Error Handling (Tidbits)
- Example 14 Defer (Tidbits)
- Example 15 Strings and Byte Arrays (Tidbits)
- Example 16 Function Type (Tidbits)

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
- Which tool do you use for package-management
- Remember: the tight relationship between package names and your directory structure (not just within a project, but within the entire workspace). What is the difference between the workspace and the project?
- How to handle the dependency `github.com/mattn/go-sqlite3`
- What is a singleton? (see Example 13b) How to run Example 13b?
- Example 14a didn't return the error :/ 
- `go fmt ./...` --> can't load package: package learning-go: no Go files in /Users/vasileiosmanolas/go/src/learning-go (?)
- This is necessary because strings are **immutable**? Do not change? Example15
- If you iterate over a string using `range`, you’ll get `runes`, not bytes (?)
- Example 16 - Using functions like this can help decouple code from specific implementations much like we achieve with interfaces. - ??