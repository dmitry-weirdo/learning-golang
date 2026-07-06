package main

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(v T) { // need to use pointer to modify the Stack, else it will be a copy
	s.data = append(s.data, v) // append to end
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) <= 0 {
		var zeroValue T
		return zeroValue, false
	}

	lastIndex := len(s.data) - 1
	value := s.data[lastIndex]
	s.data = s.data[0:lastIndex] // remove the last element

	return value, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) <= 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func isValid(s string) bool {
	var st Stack[rune]

	for _, c := range s {
		switch c {
		// the opening brackets are just pushed to the stack
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			st.Push(c)

		case ')':
			if !testMatchingBracket(&st, c, '(') {
				return false
			}

		case '}':
			if !testMatchingBracket(&st, c, '{') {
				return false
			}

		case ']':
			if !testMatchingBracket(&st, c, '[') {
				return false
			}
		}
	}

	if !st.IsEmpty() {
		fmt.Printf("After processing the string \"%v\", the stack is not empty: %v. Returning false. \n", s, string(st.data))

		return false
	}

	return true
}

func testMatchingBracket(st *Stack[rune], closingBracket rune, expectedOpeningBracket rune) bool {
	v, ok := st.Pop()

	if !ok {
		fmt.Printf("Stack is empty. No matching opening '%c' for closing '%c' found. Returning false. \n", expectedOpeningBracket, closingBracket)
		return false
	}

	if v != expectedOpeningBracket {
		fmt.Printf("Last bracket '%c' does not match the opening '%c' for closing '%c'. Returning false. \n", v, expectedOpeningBracket, closingBracket)
		return false
	}

	return true
}

func test(s string, expectedResult bool) {
	fmt.Println()
	fmt.Println("==========================")

	fmt.Printf("String: %v  \n", s)

	result := isValid(s)

	fmt.Printf("Result: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	s := "[]"
	expected := true

	test(s, expected)
}

func test2() {
	s := "([{}])"
	expected := true

	test(s, expected)
}

func test3() {
	s := "[(])"
	expected := false

	test(s, expected)
}

func test4() {
	s := "[()" // close-open match, but stack won't be empty
	expected := false

	test(s, expected)
}

func main() {
	test1()
	test2()
	test3()
	test4()
}
