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
}
