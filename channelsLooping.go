package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			message := i
			fmt.Printf("Sending message: %v \n", message)
			ch <- message
		}

		// if we don't close the channel, the program will fail with a deadlock
		close(ch)

		// panic: send on closed channel
		ch <- 666
	}()

	for message := range ch {
		fmt.Printf("Message received: %v \n", message)
	}
}
