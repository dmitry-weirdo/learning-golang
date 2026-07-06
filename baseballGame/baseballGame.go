package main

import (
	"fmt"
	"strconv"
)

type Stack struct { // we're just storing the int values
	v []int
}

func (s *Stack) push(value int) {
	s.v = append(s.v, value)
}

func (s *Stack) isEmpty() bool {
	return len(s.v) <= 0
}

func (s *Stack) top() int { // does not remove the last value
	// todo: not handling empty stack
	return s.v[len(s.v)-1]
}

func (s *Stack) topTwo() (int, int) { // does not remove the last value
	// todo: not handling stack with len < 2
	return s.v[len(s.v)-1], s.v[len(s.v)-2]
}

func (s *Stack) pop() int {
	// todo: not handling empty stack
	lastIndex := len(s.v) - 1

	lastValue := s.v[lastIndex]
	s.v = s.v[0:lastIndex] // remove last element, 2nd argument is non-inclusive

	return lastValue
}

func (s *Stack) sum() int {
	sum := 0

	for _, val := range s.v {
		sum += val
	}

	return sum
}

func calPoints(operations []string) int {
	var s Stack

	for _, op := range operations {
		switch op {
		case "+":
			v1, v2 := s.topTwo()
			sum := v1 + v2
			s.push(sum)
			fmt.Printf("Added the sum value %v = %v + %v to the stack. \n", sum, v1, v2)

		case "D":
			v := s.top()
			doubleValue := 2 * v
			s.push(doubleValue)
			fmt.Printf("Added the double value %v of value %v to the stack. \n", doubleValue, v)

		case "C":
			v := s.pop()
			fmt.Printf("Removed value %v from the stack. \n", v)

		default:
			v, _ := strconv.Atoi(op) // we don't care about the int conversion errors
			s.push(v)
			fmt.Printf("Pushed value %v to the stack. \n", v)
		}
	}

	fmt.Printf("Stack before calculating the sum: %v \n", s.v)

	return s.sum()
}

func test(operations []string, expectedResult int) {
	fmt.Println()
	fmt.Println("==========================")

	fmt.Printf("Operations: %v  \n", operations)

	result := calPoints(operations)

	fmt.Printf("Result: %v  \n", result)
	fmt.Printf("Expected result: %v  \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	operations := []string{"1", "2", "+", "C", "5", "D"}
	expected := 18

	test(operations, expected)
}

func test2() {
	operations := []string{"5", "D", "+", "C"}
	expected := 15

	test(operations, expected)
}

func main() {
	test1()
	test2()
}
