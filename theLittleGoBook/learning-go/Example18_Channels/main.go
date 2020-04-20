// Example 18a
package main

import (
	"crypto/rand"
	"fmt"
	"time"
)

// Consider a system with incoming data that we want to handle in separate goroutines. This is a common requirement.
// If we did our data-intensive processing on the goroutine which accepts the incoming data, we’d risk timing out clients.
// First, we’ll write our worker. This could be a simple function, but I’ll make it part of a structure since we haven’t seengoroutines used like this before:
type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Println("Worker %d got %d\n", w.id, data)
	}
}

// Our worker waits until data is available, then “processes” it. Dutifully, it does this in a loop, forever waiting for more data to process.
// To use this, the first thing we'd do is start some workers

func main() {
	c := make(chan int)
	for i := 0; i < 4; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}

	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * 50)
	}
}

// We don’t know which worker is going to get what data.
// What we do know, what Go guarantees, is that the data wesend to a channel will only be received by a single receiver.
// Notice that the only shared state is the channel, which we can safely receive from and send to concurrently.
// Channels provide all of the synchronization code we need and also ensure that, at any given time, only one goroutine has access to a specific piece of data.

// Example 18b - buffered channels
// what happens if we have more data coming in than we can handle?
// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"time"
// )

// type Worker struct {
// 	id int
// }

// func (w *Worker) process(c chan int) {
// 	for {
// 		data := <-c
// 		fmt.Println("Worker %d got %d\n", w.id, data)
// 		// we can simulate this by changing the worker to sleep after it has received data:
// 		time.Sleep(time.Millisecond * 500)
// 		// What’s happening is that our main code, the one that accepts the user’s incoming data (which we just simulated with a random number generator) is blocking as it sends to the channel because no receiver is available.
// 	}
// }

// func main() {
// 	c := make(chan int, 100) // we give our channel a length: 100
// 	for i := 0; i < 4; i++ {
// 		worker := &Worker{id: i}
// 		go worker.process(c)
// 	}

// 	for {
// 		c <- rand.Int()
// 		// we can get a sense that the buffered channel is, in fact, buffering by looking at the channel’s `len`:
// 		fmt.Println(len(c))
// 		// You can see that it grows and grows until it fills up, at which point sending to our channel start to block again.
// 		time.Sleep(time.Millisecond * 50)
// 	}
// }

// Example 18c - select
// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"time"
// )

// type Worker struct {
// 	id int
// }

// func (w *Worker) process(c chan int) {
// 	for {
// 		data := <-c
// 		fmt.Println("Worker %d got %d\n", w.id, data)
// 		// we can simulate this by changing the worker to sleep after it has received data:
// 		time.Sleep(time.Millisecond * 500)
// 		// What’s happening is that our main code, the one that accepts the user’s incoming data (which we just simulated with a random number generator) is blocking as it sends to the channel because no receiver is available.
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 4; i++ {
// 		worker := &Worker{id: i}
// 		go worker.process(c)
// 	}

// 	for {
// 		select {
// 		case c <- rand.Int():
// 			// optional code here
// 		default:
// 			// this can be left empty to silently drop the data
// 			fmt.Println("dropped")
// 		}
// 		time.Sleep(time.Millisecond * 50)
// 	}
// }

// Example 18d - timeout
// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"time"
// )

// type Worker struct {
// 	id int
// }

// func (w *Worker) process(c chan int) {
// 	for {
// 		select {
// 		case data := <-c:
// 			fmt.Println("Worker %d got %d\n", w.id, data)
// 		case <-time.After(time.Millisecond * 10):
// 			fmt.Println("Break time")
// 			time.Sleep(time.Second)
// 		}
// 	}
// }

// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 4; i++ {
// 		worker := &Worker{id: i}
// 		go worker.process(c)
// 	}

// 	for {
// 		select {
// 		// Pay close attention to our `select`. Notice that we’re **sending to `c`** but **receiving from `time.After`**.
// 		case c <- rand.Int():
// 			// optional code here
// 		case <-time.After(time.Millisecond * 100):
// 			// `time.After` returns a channel, so we can `select` from it.
// 			// the channel is written to after the specified time expires.
// 			fmt.Println("timed out")
// 		case t := <-time.After(time.Millisecond * 200):
// 			fmt.Println("timed out at ", t)
// 			// what happens if the `default` is added back? (remember that `default` fires immediately if no channel is available)
// 			// default:
// 			// 	// this can be left empty to silently drop the data
// 			// 	fmt.Println("dropped")
// 		}
// 		time.Sleep(time.Millisecond * 50)
// 	}
// }

// // If you’re curious, here’s what an implementation of after could look like:
// func after(d time.Duration) chan bool {
// 	c := make(chan bool)
// 	go func() {
// 		time.Sleep(d)
// 		c <- true
// 	}()
// 	return c
// }
