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
  - **Concurrent programming** demands considerably more attention and care. **Goroutines** effectively **abstract** what’s needed to run concurrent code. **Channels** help eliminate some serious bugs that can happen when data is shared by **eliminating the sharing of data**. This doesn’t just eliminate bugs, but it changes how one approaches concurrent programming. You start to **think about concurrency with respect to message passing**,rather than dangerous areas of code.Having said that, I still make extensive use of the various **synchronization primitives** found in the `sync` and `sync/atomic` **packages**. I think it’s important to be comfortable with both. I encourage you to **first focus on channels**, but when you see a simple example that needs a **short-lived lock**, consider using a `mutex` or `read-write mutex`.
  - Goroutines: similar to a thread, but it is scheduled by Go, not the OS. Code that runs in a goroutine can run concurrently with other code.
    - How we start a goroutine: using the `go` keyword followed by the function we want to execute. 
    - Multiple goroutines will end up running on the same underlying OS thread. This is often called an `M:N threading model` because we have M application threads (goroutines) running on N OS threads. The result is that a goroutine has a fraction of overhead (a few KB) than OS threads.
    - On modern hardware, it's possible to have millions of goroutines.
    - The complexity of mapping and scheduling is hidden. We just say this code should run concurrently and let Go worry about making it happen.
  - Synchronization
    - Concurrent code needs to be coordinated. 
    - Writing concurrent code requires that you pay specific attention to where and how you read and write values. In some ways, it's like programming without a garbage collector.
    - **The only concurrent thing you can safely do to a variable is to read from it. You can have as many readers as you want,but writes need to be synchronized.** 
      - There are various ways to do this, including using some truly atomic operations that rely on special CPU instructions. However, the most common approach is to use a `mutex`. A mutex serializes access to the code under lock.
    - There's a whole class of serious **bugs** that can arise when doing concurrent programming:
      1. It isn't always so obvious **what code needs to be protected**. While it might be tempting to use `coarse locks` (locks that cover a large amount of code), that undermines the very reason we’re doing concurrent programming in the first place. We generally want `fine locks`; else, we end up with a ten-lane highway thatsuddenly turns into a one-lane road.
      2. The other problem has to do with **deadlocks**. With a single lock, this isn’t a problem, but if you’re using two or morelocks around the same code, *it’s dangerously easy to have situations where goroutineA holds lockA but needs access to lockB, while goroutineB holds lockB but needs access to lockA*. It actually is possible to *deadlock with a single lock, if we forget to release it*. This isn’t as dangerous as a multi-lockdeadlock because those are really tough to spot.
    - There’s another common `mutex` called a **read-write mutex**. This exposes **two locking functions**: one to lock for reading and one to lock for writing.This distinction *allows multiple simultaneous readers while ensuring that writing is exclusive*. In Go, `sync.RWMutex` is such a lock. In addition to the `Lock`and `Unlock` methods of a `sync.Mutex`, it also exposes `RLock` and `RUnlock` methods; where `R` stands for `Read`. While read-write mutexes are commonly used, they place an additional burden on developers: we must now pay attention to not only when we’re accessing data, but also how.
    - *Part of concurrent programming isn’t so much about serializing access across the narrowest possible piece of code*; it’s also about **coordinating multiple goroutines**. For example, sleeping for 10 milliseconds isn’t a particularly elegant solution. What if a goroutine takes more than 10 milliseconds? What if it takes less and we’re just wasting cycles? Also, what if instead of just waiting for goroutines to finish, we want to tell one "hey, I have new data for you to process!"?
- Channels
  - The challenge with concurrent programming stems from **sharing data**. If your goroutines share no data, you needn’t worry about synchronizing them.
  - Channels help make concurrent programming saner by taking shared data out of the picture. **A channel is a communication pipe between goroutines which is used to pass data**. In other words, a goroutine that has data can pass it to another goroutine via a channel. The **result** is that, **at any point in time, only one goroutine has access to the data**.
  - A channel, like everything else, has a type. This is the type of data that we’ll be passing through our channel. 
  ```
  c := make(chan int)
  func worker(c chan int) { ... }
  ```
  - Receiving and sending data from and to a channel is **blocking**: *when we receive from a channel, execution of the goroutine won't continue until data is available*. Similarly, *when we send to a channel, execution won't continue until the data is received*. 
  ```
  CHANNEL <- DATA // sending to a channel
  VAR := <-CHANNEL // receiving from a channel
  ```
  - **Channels provide all of the synchronization code** we need and also ensure that, at any given time, only one goroutine has access to a specific piece of data.
- Buffered Channels
  - In cases where you need high **guarantees that the data is being processed**, you probably will want to start blocking the client. In other cases, you might be willing to loosen those guarantees. There are a few popular **strategies to do this**. The first is to **buffer the data**. 
    - *If no worker is available, we want to temporarily store the data in some sort of queue*. Channels have this buffering capability built-in.
  - Buffered channels *don’t add more capacity*; they merely **provide a queue for pending work** and a good *way to deal with a sudden spike*.
- Select
  - Even with buffering, there comes a point where we need to start dropping messages. We can’t use up an infinite amount of memory hoping a worker frees up.
  - Syntactically, select *looks a bit like a* `switch`. With it, we can provide code for when the channel isn’t available to send to.
  - **A main purpose of select is to manage multiple channels**. Given multiple channels, `select` will block until the first one becomes available. If no channel is available, `default` is executed if one is provided. A channel is randomly picked when multiple are available.
  - `select` works the same regardless of whether we’re receiving from, sending to, or any combination of channels:
    - The first available channel is chosen.
    - If multiple channels are available, one is randomly picked.
    - If no channel is available, the `default` case is executed.
    - If there’s no default, `select` **blocks**.
- Timeout
  - We’ve looked at **buffering messages** as well as simply **dropping them**. Another popular option is to **timeout**. We’re willing to **block for some time**, *but not forever*.
  - To block for a maximum amount of time, we can use the `time.After` function. `time.After` returns a **channel**, so we can `select` from it.

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
- Example 17 Concurrency
- Example 18 Channels

### Questions
- ~~What are the ENV variables in ~/.profile ?~~ 
  - `env`
- `man ssh`
  - AWS free vm OR locally from one laptop to the other
- ~~Why `Example1/main.go/Example 1b` prints the correct result when using `go build main.go`, `./main` commands ?~~
  - ~~Notice that the code compiles, there's just no entry point to run it. This is perfectly normal when you are, for example, building a library (?)~~
- Might have been upgraded but couldn't use `godoc -http=:6000` (localhost:6000)
- How to check which editor I use as default in vscode?
  - ~~vim VS bash~~ setup `vim` editor
- ~~In a terminal, I can type "vi(m)" to open Vi(m) editor. I can also type "bash" to open Bash editor. Which one is the default?~~
- ~~Would it make sense to start using vim editor straight away?~~
- ~~What is composition and what is the "battle cry" about it over inheritance (OOP)?~~~ (check **Example7**)
  - Look out for embedded structs
- Functions on structures: 
  - func (s *Saiyan) Super() {	s.Power += 10000  } VS func Super(s *Saiyan) {  s.Power += 10000  } ?
- What is the factory pattern (example 6)? And How it shields the rest of the code from worrying about allocation details?
  - Look up [Go Patterns](https://github.com/tmrts/go-patterns)
  - How to use Interfaces
    - Example 12b: see `func implementInterface(e Exister) {...}`
    - Example 12c: Empty Interface
    - 
- ~~See ?? in Example7/main.go~~
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
- Example 17c - how are two specifically goroutines writing to the same variable?
- Is there a smart way to handle goroutines? e.g. sync.RWMutex Is it with TDD ? Because we add complexity in the code that we as humans cannot track/predict its behavior in all cases. (Channels?)
- Example 18a,b - rand.Int() ? 
- Example 18c - We’re pushing out 20 messages per second, but our workers can only handle 10 per second; thus, half the messagesget dropped. (??)