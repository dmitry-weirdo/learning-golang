// package monotonicStack
package main

import (
	"container/list"
	"fmt"
)

/*
Common logic:

Direction:
- If we're searching for next: right -> left
- If we're searching for previous: left -> right

Stack:
- If we're searching for > or >=: increasing from top to bottom
- If we're searching for < or <=: decreasing from top to bottom

Removal from stack:
We're removing the opposite of what we're searching for:
- If we're searching for greater: <= current value
- If we're searching for greater or equal: < current value
- If we're searching for smaller: >= current value
- If we're searching for smaller or equal: > current value

Select top of stack into result if:
We're selecting exactly by our search clause

Push current value to stack:
- If we're searching non-inclusive (> or <): always
- If we're searching non-inclusive (>= or <=): if top of the stack != current value (to not push duplicates to the stack)
*/

func GetNextGreater(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) <= v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) > v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

		pushToStack(stack, v)
	}

	return result
}

func GetNextGreaterOrEqual(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: < current value
	// select top as result: if >= current value
	// push current value to stack: if top of the stack != current value
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) < v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) >= v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

		if (stack.Len() == 0) || (getStackTop(stack) != v) {
			pushToStack(stack, v)
		}
	}

	return result
}

func GetNextSmaller(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: decreasing from top to bottom
	// removal from stack: >= current value
	// select top as result: if < current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) >= v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) < v) {
			result[i] = getStackTop(stack)
		} else { // no next smaller element
			result[i] = noElementValue // should default to -1
		}

		pushToStack(stack, v)
	}

	return result
}

func GetNextSmallerOrEqual(a []int, noElementValue int) []int {
	// direction: right -> left
	// stack: decreasing from top to bottom
	// removal from stack: > current value
	// select top as result: if <= current value
	// push current value to stack: if top of the stack != current value
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) > v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) <= v) {
			result[i] = getStackTop(stack)
		} else { // no next smaller element
			result[i] = noElementValue // should default to -1
		}

		if (stack.Len() == 0) || (getStackTop(stack) != v) {
			pushToStack(stack, v)
		}
	}

	return result
}

func GetPrevGreater(a []int, noElementValue int) []int { // same as GetNextGreater, but left-to-right
	// direction: left -> right
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) <= v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) > v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

		pushToStack(stack, v)
	}

	return result
}

func GetPrevGreaterOrEqual(a []int, noElementValue int) []int { // same as GetNextGreaterOrEqual, but left-to-right
	// direction: left -> right
	// stack: increasing from top to bottom
	// removal from stack: < current value
	// select top as result: if >= current value
	// push current value to stack: if top of the stack != current value
	stack := list.New()

	n := len(a)
	result := make([]int, n)

	for i := 0; i < n; i++ {
		v := a[i]

		for (stack.Len() > 0) && (getStackTop(stack) < v) {
			removeFromStack(stack)
		}

		if (stack.Len() > 0) && (getStackTop(stack) >= v) {
			result[i] = getStackTop(stack)
		} else { // no next greater element
			result[i] = noElementValue // should default to -1
		}

		if (stack.Len() == 0) || (getStackTop(stack) != v) {
			pushToStack(stack, v)
		}
	}

	return result
}

// ======== helper stack functions ====== //
func pushToStack(stack *list.List, v int) { // pushes to the end of the stack
	stack.PushFront(v)
}

func removeFromStack(stack *list.List) int { // removes from the top of the stack, only called when stack is not empty
	return stack.Remove(stack.Front()).(int)
}

func getStackTop(stack *list.List) int { // only called when stack is not empty
	return stack.Front().Value.(int)
}

func printStack(l *list.List) {
	fmt.Printf("[ ")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}

	fmt.Printf("]")

	fmt.Println()
}

// ======== generic test functions ====== //
// todo: it's better to rewrite the tests as valid tests, see stack_test.go
func testGeneric(
	methodName string,
	f func([]int, int) []int,
	a []int,
	noElementValue int,
	expectedResult []int,
) {
	fmt.Println()
	fmt.Println("========================")
	fmt.Printf("Testing method \"%v\"... \n", methodName)

	fmt.Printf("Array: %v \n", a)
	fmt.Printf("No element value: %v \n", noElementValue)

	result := f(a, noElementValue)

	fmt.Printf("Result: %v \n", result)
	fmt.Printf("Expected result: %v \n", expectedResult)

	if len(result) != len(expectedResult) {
		fmt.Printf("FAILURE: expected result length = %v, actual result length = %v \n", len(expectedResult), len(result))
		return
	}

	for i, v := range result {
		if v != expectedResult[i] {
			fmt.Printf("FAILURE: expected result[%v] = %v, actual result[%v] = %v \n", i, expectedResult[i], i, v)
			return
		}
	}
}

func testGetNextGreater(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetNextGreater", GetNextGreater, a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetNextGreaterOrEqual", GetNextGreaterOrEqual, a, noElementValue, expectedResult)
}

func testGetNextSmaller(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetNextSmaller", GetNextSmaller, a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetNextSmallerOrEqual", GetNextSmallerOrEqual, a, noElementValue, expectedResult)
}

func testGetPrevGreater(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetPrevGreater", GetPrevGreater, a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual(a []int, noElementValue int, expectedResult []int) {
	testGeneric("GetPrevGreaterOrEqual", GetPrevGreaterOrEqual, a, noElementValue, expectedResult)
}

// ======== testGetNextGreater ====== //
func testGetNextGreater1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{2, 3, 4, 5, -1}

	testGetNextGreater(a, noElementValue, expectedResult)
}

func testGetNextGreater2() {
	a := []int{1, 3, 4, 2}
	noElementValue := -1
	expectedResult := []int{3, 4, -1, -1}

	testGetNextGreater(a, noElementValue, expectedResult)
}

func testGetNextGreater3() {
	a := []int{1, 1, 1, 2}
	noElementValue := -1
	expectedResult := []int{2, 2, 2, -1}

	testGetNextGreater(a, noElementValue, expectedResult)
}

func testGetNextGreater4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1}

	testGetNextGreater(a, noElementValue, expectedResult)
}

func testGetNextGreaterSuite() {
	testGetNextGreater1()
	testGetNextGreater2()
	testGetNextGreater3()
	testGetNextGreater4()
}

// ======== testGetNextGreaterOrEqual ====== //
func testGetNextGreaterOrEqual1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{2, 3, 4, 5, -1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual2() {
	a := []int{1}
	noElementValue := -1
	expectedResult := []int{-1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual3() {
	a := []int{10, 1, 1, 6}
	noElementValue := -1
	expectedResult := []int{-1, 1, 6, -1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{1, 1, 1, -1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual5() {
	a := []int{5, 4, 3, 2, 1}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1, -1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqual6() {
	a := []int{5, 5, 4, 6, 9, 1}
	noElementValue := -1
	expectedResult := []int{5, 6, 6, 9, -1, -1}

	testGetNextGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetNextGreaterOrEqualSuite() {
	testGetNextGreaterOrEqual1()
	testGetNextGreaterOrEqual2()
	testGetNextGreaterOrEqual3()
	testGetNextGreaterOrEqual4()
	testGetNextGreaterOrEqual5()
	testGetNextGreaterOrEqual6()
}

// ======== testGetNextSmaller ====== //
func testGetNextSmaller1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1, -1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmaller2() {
	a := []int{1}
	noElementValue := -1
	expectedResult := []int{-1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmaller3() {
	a := []int{10, 1, 1, 6}
	noElementValue := -1
	expectedResult := []int{1, -1, -1, -1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmaller4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmaller5() {
	a := []int{5, 4, 3, 2, 1}
	noElementValue := -1
	expectedResult := []int{4, 3, 2, 1, -1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmaller6() {
	a := []int{5, 5, 4, 6, 9, 1}
	noElementValue := -1
	expectedResult := []int{4, 4, 1, 1, 1, -1}

	testGetNextSmaller(a, noElementValue, expectedResult)
}

func testGetNextSmallerSuite() {
	testGetNextSmaller1()
	testGetNextSmaller2()
	testGetNextSmaller3()
	testGetNextSmaller4()
	testGetNextSmaller5()
	testGetNextSmaller6()
}

// ======== testGetNextSmallerOrEqual ====== //
func testGetNextSmallerOrEqual1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1, -1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual2() {
	a := []int{1}
	noElementValue := -1
	expectedResult := []int{-1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual3() {
	a := []int{10, 1, 1, 6}
	noElementValue := -1
	expectedResult := []int{1, 1, -1, -1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{1, 1, 1, -1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual5() {
	a := []int{5, 4, 3, 2, 1}
	noElementValue := -1
	expectedResult := []int{4, 3, 2, 1, -1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqual6() {
	a := []int{5, 5, 4, 6, 9, 1}
	noElementValue := -1
	expectedResult := []int{5, 4, 1, 1, 1, -1}

	testGetNextSmallerOrEqual(a, noElementValue, expectedResult)
}

func testGetNextSmallerOrEqualSuite() {
	testGetNextSmallerOrEqual1()
	testGetNextSmallerOrEqual2()
	testGetNextSmallerOrEqual3()
	testGetNextSmallerOrEqual4()
	testGetNextSmallerOrEqual5()
	testGetNextSmallerOrEqual6()
}

// ======== testGetPrevGreater ====== //
func testGetPrevGreater1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1, -1}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreater2() {
	a := []int{1}
	noElementValue := -1
	expectedResult := []int{-1}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreater3() {
	a := []int{10, 1, 1, 6}
	noElementValue := -1
	expectedResult := []int{-1, 10, 10, 10}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreater4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreater5() {
	a := []int{5, 4, 3, 2, 1}
	noElementValue := -1
	expectedResult := []int{-1, 5, 4, 3, 2}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreater6() {
	a := []int{5, 5, 4, 6, 9, 1}
	noElementValue := -1
	expectedResult := []int{-1, -1, 5, -1, -1, 9}

	testGetPrevGreater(a, noElementValue, expectedResult)
}

func testGetPrevGreaterSuite() {
	testGetPrevGreater1()
	testGetPrevGreater2()
	testGetPrevGreater3()
	testGetPrevGreater4()
	testGetPrevGreater5()
	testGetPrevGreater6()
}

// ======== testGetPrevGreaterOrEqual ====== //
func testGetPrevGreaterOrEqual1() {
	a := []int{1, 2, 3, 4, 5}
	noElementValue := -1
	expectedResult := []int{-1, -1, -1, -1, -1}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual2() {
	a := []int{1}
	noElementValue := -1
	expectedResult := []int{-1}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual3() {
	a := []int{10, 1, 1, 6}
	noElementValue := -1
	expectedResult := []int{-1, 10, 1, 10}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual4() {
	a := []int{1, 1, 1, 1}
	noElementValue := -1
	expectedResult := []int{-1, 1, 1, 1}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual5() {
	a := []int{5, 4, 3, 2, 1}
	noElementValue := -1
	expectedResult := []int{-1, 5, 4, 3, 2}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqual6() {
	a := []int{5, 5, 4, 6, 9, 1}
	noElementValue := -1
	expectedResult := []int{-1, 5, 5, -1, -1, 9}

	testGetPrevGreaterOrEqual(a, noElementValue, expectedResult)
}

func testGetPrevGreaterOrEqualSuite() {
	testGetPrevGreaterOrEqual1()
	testGetPrevGreaterOrEqual2()
	testGetPrevGreaterOrEqual3()
	testGetPrevGreaterOrEqual4()
	testGetPrevGreaterOrEqual5()
	testGetPrevGreaterOrEqual6()
}

func main() {
	testGetNextGreaterSuite()
	testGetNextGreaterOrEqualSuite()
	testGetNextSmallerSuite()
	testGetNextSmallerOrEqualSuite()

	testGetPrevGreaterSuite()
	testGetPrevGreaterOrEqualSuite()
}
