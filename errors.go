package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is an error")
	fmt.Println(err)

	// freaking Go doesn't like an error to end with a dot or start with a capital letter
	// %w can handle another Error (i.e. like exception chain)
	err2 := fmt.Errorf("This error wraps the first one: %w", err)
	fmt.Println(err2)
}
