package main

import (
	"fmt"
	"time"
)

func main() {
	// ch MUST be created with make
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch1 <- "message to channel 2"
	}()

	time.Sleep(10 * time.Millisecond) // takes Duration

	// in case of both cases can execute, the executed is selected randomly
	select {
	// receive from channel 1
	case message := <-ch1:
		fmt.Println(message)

	// receive from channel 2
	case message := <-ch2:
		fmt.Println(message)

	default: // make it non-blocking
		fmt.Println("no message available")
	}

	// deadlock if nothing happens in the channels
	// fatal error: all goroutines are asleep - deadlock!
}
