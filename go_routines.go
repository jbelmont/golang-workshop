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
