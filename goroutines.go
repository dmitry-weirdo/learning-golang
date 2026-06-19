package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // do this BEFORE declaring the concurrent task

	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done() // decreases counter by 1 (Add -1)
	}() // invoke immediately

	fmt.Println("This happens synchronous")

	wg.Wait() // waits until counter is 0
}
