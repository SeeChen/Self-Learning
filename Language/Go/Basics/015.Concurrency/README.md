<div align=center>

# Chapter 15: Concurrency in Go

[Summary](#1-summary)</br>
[Technology](#2-technology)</br>
[Key Concepts](#3-key-concepts)</br>
[Run & Output](#4-run--output)</br>
[Practice](#5-practice)</br>
[References](#6-references)

</div>

## 1. Summary
This chapter introduces **concurrency in Go**, focusing on:
    - **Goroutines**: Lightweight threads managed by Go runtime.
    - **Channels**: Communication mechanism for goroutines.
    - **Buffered Channels**: Allow asynchronous communication.
    - **GMP Model**: How Go schedules goroutines internally.

It demonstrates:
    - Running lightweight concurrent tasks with Goroutines
    - Using Channels for safe communication between Goroutines
    - Buffered vs. unbuffered channels
    - Closing channels and iterating with `range`
    - Go’s concurrency model (G: Goroutine, M: Machine, P: Processor)

## 2. Technology
- **Language**: Go 1.24  
- **Packages Used**:
    - `fmt`: for formatted printing
    - `time`: for simulating delays and synchronization  
- **Features Covered**:
    - Goroutines (`go` keyword)
    - Channels (unbuffered and buffered)
    - Channel direction (`<-`)
    - Closing channels
    - GMP Scheduler model (theory)

## 3. Key Concepts
- **Goroutine**:
    - Lightweight thread managed by Go runtime  
    - Created using `go func()` syntax  
    - Non-blocking, can run thousands efficiently  

- **Channel**:
    - Communication mechanism between Goroutines  
    - Created with `make(chan type)`  
    - Send: `ch <- value`, Receive: `val := <-ch`  
    - Can be **buffered** (`make(chan int, size)`) or **unbuffered**  

- **Buffered Channels**:
    - Allow async communication until buffer is full  
    - Once full, senders block until space is available  

- **Closing Channels**:
    - Done with `close(ch)`  
    - Important for avoiding deadlocks in `range` iteration  

- **Scheduler (GMP model)**:
    - **G**: Goroutines  
    - **M**: OS threads  
    - **P**: Logical processors managing Goroutine execution  

## 4. Run & Output
Run the program using:
```bash
go run main.go
```

Expected output:
```bash
Main: [0] First.
Goroutine: [0] First.
Goroutine: [1] Second.
Main: [1] Second.
Goroutine: [2] Third.
Main: [2] Third.
Goroutine: [3] SeeChen.
Main: [3] SeeChen.
Main: [4] Lee.
Goroutine: [4] Lee.

6 + 13 = 19
Received: 0
Sent: 10
Sent: 9
Received: 1
Sent: 8
Sent: 7
Sent: 6
Received: 10
Sent: 5
Received: 9
Sent: 4
Received: 8
Sent: 3
Received: 7
Sent: 2
Received: 6
Sent: 1
Received: 5
Received: 4
Received: 3
Received: 2
Received: 1
```

## 5. Practice
1. Modify the program to run **two different Goroutines** that process strings and numbers at the same time.  
2. Implement a **buffered channel** with capacity `3` and observe how sending/receiving changes.  
3. Experiment with removing `close(ch)` — what happens to the loop?  
4. Add a **select statement** to receive from multiple channels.  
5. Simulate a **producer-consumer** model using Goroutines and buffered channels. 

## 6. References
1. [Go Official Documentation](https://go.dev/doc/)  
2. [Go by Example - Goroutines](https://gobyexample.com/goroutines)  
3. [Go by Example - Channels](https://gobyexample.com/channels)  
4. [Go by Example - Channel Buffering](https://gobyexample.com/channel-buffering)  
5. [Spec: Concurrency](https://go.dev/ref/spec#Concurrency)  

---
<div align="right">

###### *Last Modified by [SeeChen](https://github.com/SeeChen/) @ 01-SEPT-2025 16:52 UTC +08:00*
</div>