package main

import (
	"fmt"
	"sync"
)

func main() {
	defer command.Parse().Stop()
	// Create a waitgroup
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		// Add to the wait group **before** creating the
		//goroutine
		wg.Add(1)
		go func() {
			fmt.Println("Execute routine ", i)
			// Make sure the waitgroup knows about
			// goroutine completion
			defer wg.Done()
			// Do work
		}()
	}
	// Wait until all goroutines are done
	wg.Wait()
	fmt.Println("Release waitgroup")
}
