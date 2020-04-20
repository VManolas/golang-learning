// Example 17a
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go process()
	time.Sleep(time.Millisecond * 10) // this is bad, don't do this! :P
	fmt.Println("done")
}

func process() {
	fmt.Println("processing")
}

// Example 17b - using an anonymous function
// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	fmt.Println("start")
// 	go func() {
// 		fmt.Println("processing")
// 	}()
// 	time.Sleep(time.Millisecond * 10) // this is bad, don't do this! :P
// 	fmt.Println("done")
// }

// we had to `Sleep` for a few milliseconds because the main process exits before the goroutine gets a chance to execute
// the process doesn’t wait until all goroutines are finished before exiting. To solve this, we need to coordinate our code.

// Example 17c - synchronization
// package main

// import (
// 	"fmt"
// 	"time"
// )

// var counter = 0

// func main() {
// 	for i := 0; i < 20; i++ {
// 		go incr()
// 	}
// 	time.Sleep(time.Millisecond * 10)
// }

// func incr() {
// 	counter++
// 	fmt.Println(counter)
// }

// the behavior is undefined because we potentially have multiple (two in this case) goroutines writing to the same variable, counter, at the same time.
// Or, just as bad, one goroutine would be reading counter while another writes to it.

// Example 17d - mutex
// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	counter = 0
// 	lock    sync.Mutex
// )

// func main() {
// 	for i := 0; i < 20; i++ {
// 		go incr()
// 	}
// 	time.Sleep(time.Millisecond * 10)
// }

// func incr() {
// 	lock.Lock()
// 	defer lock.Unlock()
// 	counter++
// 	fmt.Println(counter)
// }

// A mutex serializes access to the code under lock.
// The reason we simply define our lock as `lock sync.Mutex` is because the **default value** of a `sync.Mutex` is `unlocked`.

// Example 17e - deadlock
// package main

// import (
// 	"sync"
// 	"time"
// )

// var (
// 	lock sync.Mutex
// )

// func main() {
// 	go func() { lock.Lock() }()
// 	time.Sleep(time.Millisecond * 10)
// 	lock.Lock()
// }

// fatal error: all goroutines are asleep - deadlock!
