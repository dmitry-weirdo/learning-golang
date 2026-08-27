package main

import "fmt"

func canSeePersonsCount(heights []int) []int {
	// We basically do the "next greater" monotonic stack,
	// and count the values:
	// - Removed as <= values - we can see them
	// - If there is a top of stack (i.e. next greater element) - we count it as well.
	// (strictly greater since all values are distinct)
	return GetNextGreater(heights)
}

func GetNextGreater(a []int) []int {
	// direction: right -> left
	// stack: increasing from top to bottom
	// removal from stack: <= current value
	// select top as result: if > current value
	// push current value to stack: always
	stack := createStack()

	n := len(a)
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		v := a[i]

		for stackIsNotEmpty(stack) && (getStackTop(stack) <= v) {
			removeFromStack(stack)

			// we will see the next smaller from the current element
			result[i]++
		}

		if stackIsNotEmpty(stack) && (getStackTop(stack) > v) {
			// there is next greater element -> we will see it as well, count it
			result[i]++
		}

		pushToStack(stack, v)
	}

	return result
}

// ======== helper stack functions ====== //
// ======== stack of int values ====== //
func createStack() *[]int {
	stack := make([]int, 0)
	return &stack
}

func stackIsNotEmpty(stack *[]int) bool {
	return len(*stack) > 0
}

func stackIsEmpty(stack *[]int) bool {
	return len(*stack) == 0
}

func pushToStack(stack *[]int, v int) { // pushes to the end of the stack
	// we push to the end of the slice = top of the stack
	*stack = append(*stack, v)
}

func removeFromStack(stack *[]int) int { // removes from the top of the stack, only called when stack is not empty
	lastElement := (*stack)[len(*stack)-1]

	*stack = (*stack)[:len(*stack)-1] // remove the last element

	return lastElement
}

func getStackTop(stack *[]int) int { // only called when stack is not empty
	return (*stack)[len(*stack)-1]
}

// ======== tests ====== //
func test(arr []int, expectedResult []int) {
	fmt.Println()
	fmt.Println("========================")

	fmt.Printf("Heights of people: %v \n", arr)

	result := canSeePersonsCount(arr)

	fmt.Printf("Persons the i-th person can see to the right of him: %v \n", result)
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

func test1() {
	test(
		[]int{10, 6, 8, 5, 11, 9},
		[]int{3, 1, 2, 1, 1, 0},
	)
}

func test2() {
	test(
		[]int{5, 1, 2, 3, 10},
		[]int{4, 1, 1, 1, 0},
	)
}

func main() {
	// 1944. Number of Visible People in a Queue
	test1()
	test2()
}
