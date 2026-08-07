package main

import (
	"fmt"
)

func removeDuplicateLetters(s string) string {
	// collect the last occurrence of every letter
	// todo: we can use an array of 26, but it will require initializing to -1
	lastPos := make(map[byte]int)

	for i, c := range s {
		lastPos[byte(c)] = i // later positions will overwrite previous positions
	}

	// keeps elements in the stack (added to the result), but
	visited := make(map[byte]bool)

	// stack tracks last elements
	stack := createStack()

	for i, c := range s {
		ch := byte(c)

		if visited[ch] { // char already visited -> skip the duplicate
			continue
		}

		// we remove bigger characters from the stack, but only if they are present in the later positions than the current
		for stackIsNotEmpty(stack) &&
			(getStackTop(stack) > ch) && // bigger character
			(lastPos[getStackTop(stack)] > i) { // appears after the current position
			biggerCharThatAppearsLater := removeFromStack(stack)

			// this char is not visited anymore -> will be added from the later index
			visited[biggerCharThatAppearsLater] = false
		}

		pushToStack(stack, ch)
		visited[ch] = true
	}

	// stack contains earliest letters at the bottom, but we convert the underlying slice that is bottom -> top :)
	return string(*stack)
}

func createStack() *[]byte {
	stack := make([]byte, 0)
	return &stack
}

func stackIsNotEmpty(stack *[]byte) bool {
	return len(*stack) > 0
}

func stackIsEmpty(stack *[]byte) bool {
	return len(*stack) == 0
}

func pushToStack(stack *[]byte, v byte) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStack(stack *[]byte) byte { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTop(stack *[]byte) byte { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

func test(s string, expectedResult string) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("String: %v \n", s)

	result := removeDuplicateLetters(s)

	fmt.Printf("Ordered unique letters in the string: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
		return
	}
}

func test1() {
	s := "bcabc"
	expected := "abc"

	test(s, expected)
}

func test2() {
	s := "cbacdcbc"
	expected := "acdb"

	test(s, expected)
}

func main() {
	// 316. Remove Duplicate Letters
	test1()
	test2()
}
