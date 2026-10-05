package main

import "fmt"

func scoreOfParentheses(s string) int {
	return scoreOfParentheses_stack(s)
}

func scoreOfParentheses_stack(s string) int {
	var stack Stack[int]
	stack.Push(0) // The score of the current frame

	for _, ch := range s {
		if ch == '(' {
			stack.Push(0)
			continue
		}

		currentLevel, _ := stack.Pop()
		previousLevel, _ := stack.Pop()
		stack.Push(previousLevel + max(2*currentLevel, 1))
	}

	result, _ := stack.Pop()
	return result
}

func scoreOfParentheses_heuristic(s string) int {
	sum := 0

	//currentLevelSum := 0
	level := 0

	//m := make(map[int]int) // level to level score

	//var stack Stack[]

	for i, ch := range s {
		if ch == '(' {
			level++
		} else {
			//if m[level]

			level--
		}

		fmt.Printf("i: %v, s[%v] = '%c', level = %v \n", i, i, ch, level)
	}

	return sum
}

type BracketLevel struct {
	level int
	sum   int // sum on this level
}

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

func test(s string, expectedResult int) {
	fmt.Println()
	fmt.Println("====================")

	fmt.Printf("String of brackets: %v \n", s)

	result := scoreOfParentheses(s)

	fmt.Printf("Parentheses score: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	test("()", 1)
}

func test2() {
	test("(())", 2)
}

func test3() {
	test("()()", 2)
}

func test4() {
	test("(()(()))", 6)
}

func main() {
	// 856. Score of Parentheses
	test1()
	test2()
	test3()
	test4()
}
