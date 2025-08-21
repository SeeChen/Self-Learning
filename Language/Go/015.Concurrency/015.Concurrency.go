package main

import (
	"fmt"
	"time"
)

func goroutine(x []string) {
	for i, v := range x {
		fmt.Printf("Goroutine: [%d] %s.\n", i, v)
		time.Sleep(100 * time.Millisecond)
	}
}

func goChannel(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	ch <- sum // ch is a Channel that receives data.
	// In the channel, the function does not need to return a value
	// But passes the return value back to the main function throught the channel parameter.
}

func main() {

	// Golang supports concurrency and provides a simple and efficient way to achieve concurrency through Goroutines and Channels.
	// Goroutines:
	// 	- Concurrent execution unit in Golang, similar to lightweight threads.
	//	- Goroutines scheduling is managed by the Golang runtime and users do not need to manually allocate threads.
	//	- To start Goroutine, use the `go` keyword.
	// 	- Goroutines is non-blocking and can run thousands of Goroutines efficiently.

	// Chhannel:
	// 	- Go's mechanism for communication between Goroutines.
	// 	- Support synchronization and data sharing, avoiding explicit locking mechanismes
	// 	- Created by the `chan` keyword, data is sent and received throught the `<-` operator.

	// Scheduler:
	// 	- The Golang scheduler is based on the GMP model and assigns Goroutines to OS threads for execution.
	// 	- It efficiently manages concurrency through the coordination of M and P.
	//
	// G: Goroutine
	// M: Machine
	// P: Processor

	// # GOROUTINE
	// Goroutine syntax `go func(parameters)`
	var x []string = []string{"First", "Second", "Third", "SeeChen", "Lee"}
	go goroutine(x) // Excuting the goroutine
	for i, v := range x {
		fmt.Printf("Main: [%d] %v.\n", i, v)
		time.Sleep(100 * time.Millisecond)
	}

	// By executing the above code, we will see the output of two functions, but not in a fixed order.

	// # Channel
	// Channel is used to transfer data between two Goroutines.
	// We can create a Channel using the `make` function and use the `<-` operator to receive or send data.
	// ch <- value: ch is a Channel that receives data.
	// value <- ch: ch is a Channel that sends data.

	// By default, a Channel can both receive and send data, but once a direction is specified, it can only snd or receive data.
	// The channel must be created before use.

	s := []int{1, 3, 4, 5, 7, 8, -10, -4, 5}
	ch := make(chan int) // Create a channel type variable.

	go goChannel(s[:len(s)/2], ch)
	go goChannel(s[len(s)/2:], ch)

	// This channel is a receiving variable
	p, q := <-ch, <-ch

	fmt.Printf("\n%d + %d = %d\n", p, q, p+q)

	// # Channel Buffer
	// We can create a buffered Channel by specifying its buffer size as the second argument to the `make` function.
	// Buffered channels allow senders and receivers to communicate asynchronously, temporarily decoupling the two.
	// This means that data sent to the channel can be stored in the buffer, waiting to be received later.

	// However, the buffer size is limited. If the buffer is full, the sender will not be able to send any more data until data is retrieved from the buffer.

	// If you don't specify a size, the channel is unbuffered, meaning it has a buffer size of 0.
	// If the buffer is full, the channel operation will block the sendin goroutine until there is space available.

	var ch1 chan int = make(chan int, 5) // Defined a variable it's five buffer size.
	ch1 <- 0
	ch1 <- 1

	go func(ch0 chan int) {
		for i := 10; i > 0; i-- {
			ch0 <- i
			fmt.Printf("Sent: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}

		// Close the channel
		// Closing the channel is crucial.
		// If you don't close the channel, the `range` loop will block and wait for more data.
		// Causing the program to hand or enter a dealock
		close(ch0)
	}(ch1)

	// Traverse the Channel
	for val := range ch1 {
		fmt.Printf("Received: %d\n", val)
		time.Sleep(200 * time.Millisecond)
	}
}
