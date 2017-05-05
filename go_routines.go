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

func routines() {
	var wg sync.WaitGroup

	// you can also add these one at
	// a time if you need to

	wg.Add(1)

	messages := []string{
		"My", "Next", "message",
	}
	blockingPrint(messages)

	go func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
	}("Block Print")

	wg.Wait()
}
