# Go Routines

## Definition of a goroutine
* A goroutine is a lightweight thread of execution.
* A goroutine is a function that is capable of running concurrently with other functions.
* Goroutines are multiplexed onto multiple OS threads so if one should block, such as while waiting for I/O, others continue to run. 
* Goroutine design hides many of the complexities of thread creation and management.

[Goroutines Effective Go](https://golang.org/doc/effective_go.html#goroutines)

## Usage of a goroutine
To create a goroutine we use the keyword go followed by a function invocation

```go
package workshop

import (
	"fmt"
	"sync"
)

func blockingPrint(messages []string) {
	for i := 0; i < len(messages); i++ {
		fmt.Println(messages[i])
	}
}

func goRoutines() {
	var wg sync.WaitGroup
	wg.Add(2)

	messages := []string{
		"My", "Next", "message",
	}
	blockingPrint(messages)

	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("Block Print")

	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("Block Print number 2")

	wg.Wait()
}

// Output here
// My
// Next
// message
// Block Print number 2
// Block Print
```


1. Notice here that the output of `My, Next, message` are blocking calls 
    1. Meanwhile the next msgs `Block Print` and `Block Print number 2` are done concurrently with goroutines.
    1. Also notice that the last goroutine called is the first message shown so they execute in reverse order

[Go Routines Playground](https://play.golang.org/p/8HX6vc2eMB)