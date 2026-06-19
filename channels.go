package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// ch MUST be created with make
	ch := make(chan string)

	wg.Add(1) // do this BEFORE declaring the concurrent task

	go func() {
		message := "the message"
		fmt.Printf("Sending message \"%v\" to the channel. \n", message)

		ch <- message

		// !!! no Done() here, only in the last goroutine in the pipeline
		//wg.Done() // decreases counter by 1 (Add -1)

	}() // invoke immediately

	go func() {
		message := <-ch
		fmt.Printf("Received message \"%v\" from the channel. \n", message)

		// Done only at the end of the pipeline
		wg.Done()
	}()

	fmt.Println("This happens synchronous.")

	wg.Wait() // waits until counter is 0
}
