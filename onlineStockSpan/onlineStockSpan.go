package main

import "fmt"

type StockSpanner struct {
	stack          *[]MatchingElement
	result         []MatchingElement
	noElementValue int
}

func Constructor() StockSpanner {
	// init phase of GetPrevGreaterWithIndexes
	stack := createStackWithIndex()
	result := make([]MatchingElement, 0) // 0 elements yet

	return StockSpanner{
		stack:          stack,
		result:         result,
		noElementValue: -1,
	}
}

func (this *StockSpanner) Next(price int) int {
	// basically, we need the "get previous greater" index, but with adding a new value to the stack

	// todo: yes, this can be optimized -> we don't need to store the complete result array, we can just get all the info from the stack
	// This is the post-init phase of GetPrevGreaterWithIndexes.
	// Init phase was in the constructor.
	// We don't iterate the array, we handle value one-by-one

	// we execute just 1 iteration in the stack
	i := len(this.result)
	//fmt.Printf("i: %v, result: %v, stack: %v \n", i, this.result, this.stack)

	// add 1 non-filled element[i] to result
	this.result = append(this.result, MatchingElement{})

	v := price

	stack := this.stack

	for stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value <= v) {
		removeFromStackWithIndex(stack)
	}

	if stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value > v) {
		this.result[i] = getStackTopWithIndex(stack)
	} else { // no prev greater element
		this.result[i] = MatchingElement{value: this.noElementValue, index: -1} // should default to -1
	}

	currentElement := MatchingElement{value: v, index: i}
	pushToStackWithIndex(stack, currentElement)

	if this.result[i].value == this.noElementValue { // no prev greater -> all elements before should be counted
		return i + 1
	}

	// subtract i - prevGreaterIndex
	return i - this.result[i].index
}

type MatchingElement struct {
	value int
	index int // index of value in the array
}

func GetPrevGreaterWithIndexes(a []int, noElementValue int) []MatchingElement {
	// direction: left -> right
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := createStackWithIndex()

	n := len(a)
	result := make([]MatchingElement, n)

	for i := 0; i < n; i++ {
		v := a[i]

		for stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value <= v) {
			removeFromStackWithIndex(stack)
		}

		if stackIsNotEmptyWithIndex(stack) && (getStackTopWithIndex(stack).value > v) {
			result[i] = getStackTopWithIndex(stack)
		} else { // no prev greater element
			result[i] = MatchingElement{value: noElementValue, index: -1} // should default to -1
		}

		currentElement := MatchingElement{value: v, index: i}
		pushToStackWithIndex(stack, currentElement)
	}

	return result
}

// ======== stack of MatchingElement values begin ====== //
func createStackWithIndex() *[]MatchingElement {
	stack := make([]MatchingElement, 0)
	return &stack
}

func stackIsNotEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) > 0
}

func stackIsEmptyWithIndex(stack *[]MatchingElement) bool {
	return len(*stack) == 0
}

func pushToStackWithIndex(stack *[]MatchingElement, v MatchingElement) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStackWithIndex(stack *[]MatchingElement) MatchingElement { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTopWithIndex(stack *[]MatchingElement) MatchingElement { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

// ======== stack of MatchingElement values end ====== //

func testNext(s *StockSpanner, price int, expectedResult int) {
	fmt.Println()
	fmt.Printf("Executing StockSpanner.Next(%v)... \n", price)

	result := s.Next(price)

	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if result != expectedResult {
		fmt.Printf("FAILURE: expected result = %v, actual result = %v \n", expectedResult, result)
	}
}

func test1() {
	stockSpanner := Constructor()

	s := &stockSpanner

	testNext(s, 100, 1) // 100
	testNext(s, 80, 1)  // 80
	testNext(s, 60, 1)  // 60
	testNext(s, 70, 2)  // 60, 70
	testNext(s, 60, 1)  // 60
	testNext(s, 75, 4)  // 60, 70, 60, 75
	testNext(s, 85, 6)  // 80, 60, 70, 60, 75, 85
}

func test2() {
	// failing test-case 3 / 101
	stockSpanner := Constructor()

	s := &stockSpanner

	testNext(s, 50, 1) // 1
	testNext(s, 98, 2) // 50, 98
	testNext(s, 54, 1) // 54
	testNext(s, 6, 1)  // 6
	testNext(s, 34, 2) // 6, 34
}

func main() {
	// 901. Online Stock Span
	test1()
	test2()
}
